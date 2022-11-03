package main

import (
	"fmt"
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
    "github.com/amthorn/task_wizard/ent"
	pb "github.com/amthorn/task_wizard/services/project_service/src/proto"

    entsql "entgo.io/ent/dialect/sql"
)

type TWDatabase struct {
	db *ent.Client
}

func (this *TWDatabase) entToProto(model *ent.Project) (*pb.Project) {
	return &pb.Project{Id: int64(model.ID), Name: model.Name}
}

func (this *TWDatabase) getDbUri() (string) {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_PORT"), 
		os.Getenv("DB_NAME"),
	)
}

func (this *TWDatabase) setUpDB() (*ent.Client, error) {
	_db, err := sql.Open("mysql", this.getDbUri())
	if err != nil {
		log.Fatalf("Could not connect to database: %w", err)
	}

	max_connection_lifetime, err := strconv.Atoi(os.Getenv("DB_MAX_CONNECTION_LIFETIME"))
	if err != nil { panic(err) }
	_db.SetConnMaxLifetime(time.Minute * time.Duration(max_connection_lifetime))

	max_open_connections, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	if err != nil { panic(err) }
	_db.SetMaxOpenConns(max_open_connections)

	max_idle_connections, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	if err != nil { panic(err) }
	_db.SetMaxIdleConns(max_idle_connections)

	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	log.Info("Pinging db...")
	if err := _db.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %w", err)
	}
	log.Info("Database is reachable")

	this.db = ent.NewClient(ent.Driver(entsql.OpenDB("mysql", _db)))

    if err := this.db.Schema.Create(context.Background()); err != nil {
        log.Fatalf("Failed creating schema resources: %v", err)
    }
	log.Info("Migration complete")

	return this.db, nil
}

func NewTWDatabase() (*TWDatabase) {
	tw_db := &TWDatabase{}
	tw_db.setUpDB()
	return tw_db
}