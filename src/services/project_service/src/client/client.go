package main

import (
	"fmt"
	"context"
	"log"
	"time"
	"flag"

	pb "github.com/amthorn/task_wizard/services/project_service/src/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	host = flag.String("host", "localhost", "The host to send the request on")
	port = flag.Int("port", 8080, "The port to send the request on")
)

func newConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *host, *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	return conn, nil
}

func newClient(conn *grpc.ClientConn) (pb.ProjectServiceClient, error) {
	return pb.NewProjectServiceClient(conn), nil
}

func getProject() {
	conn, err := newConnection()
	if err != nil {
		log.Fatalf("Could not connect to server")
	}
	defer conn.Close()

	client, err := newClient(conn)
	if err != nil {
		log.Fatalf("Could not create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.GetProject(ctx, &pb.ProjectId{Id: 1})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
		// lets print the error code which is `INVALID_ARGUMENT`
		fmt.Println(errStatus.Code())
		// log.Fatalf("could not get project: %v", err)
	}
	log.Printf("NAME IS: %s", r.GetName())
	r, err = client.GetProject(ctx, &pb.ProjectId{Id: 7})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
		// lets print the error code which is `INVALID_ARGUMENT`
		fmt.Println(errStatus.Code())
		// log.Fatalf("could not get project: %v", err)
	}
	log.Printf("NAME IS: %s", r.GetName())
	r, err = client.CreateProject(ctx, &pb.Project{Name: "Fourth"})
	log.Printf("NAME IS: %s", r.GetName())
	log.Printf("ID IS: %d", r.GetId())

	r, err = client.GetProject(ctx, &pb.ProjectId{Id: r.GetId()})
	log.Printf("NAME IS: %s", r.GetName())
}

func main() {
	flag.Parse()
	getProject()
}
