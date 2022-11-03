package main

import (
	"fmt"
	"context"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	log "github.com/sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"

    "github.com/amthorn/task_wizard/ent"
    // "github.com/amthorn/task_wizard/ent/project"
	pb "github.com/amthorn/task_wizard/services/project_service/src/proto"

)

type Projects struct {
	db TWDatabase
}

func (this *Projects) entToProto(model *ent.Project) (*pb.Project) {
	return &pb.Project{Id: int64(model.ID), Name: model.Name}
}

func (this *Projects) GetProject(ctx context.Context, id int64) (*pb.Project, error) {
	log.Info("Looking up project")
	result, err := this.db.db.Project.
		Query().
		Where(project.ID(int(id))).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Project with id '%d' does not exist", id)
	}
	return this.entToProto(result), nil
}

func (this *Projects) CreateProject(ctx context.Context, project *pb.Project) (*pb.Project, error) {
	log.Info("Looking up project")
	// TODO(avathorn): Can set fields dynamically instead of each one being set??
	result, err := this.db.db.Project.
		Create().
		SetName(project.Name).
		Save(ctx)
	if err != nil {
		log.Info(fmt.Sprintf("Project with ID '%d' already exists", result.ID))
		return nil, fmt.Errorf("failed creating project: %w", err)
	}
	return this.entToProto(result), nil
}

func NewProjects() (*Projects) {
	projects := &Projects{}
	projects.db = *NewTWDatabase()
	return projects
}