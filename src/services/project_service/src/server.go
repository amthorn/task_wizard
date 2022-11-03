package main

import (
	"fmt"
	"context"
	"net"

	"google.golang.org/grpc"
	log "github.com/sirupsen/logrus"

	pb "github.com/amthorn/task_wizard/services/project_service/src/proto"
)

type Server struct {
	pb.UnimplementedProjectServiceServer
	projects Projects
}

func (this *Server) GetProject(ctx context.Context, req *pb.ProjectId) (*pb.Project, error) {
	log.Info("got request for ID:", req.Id)
	result, err := this.projects.GetProject(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("Failed getting project: %w", err)
	}
    log.Info("project returned: ", result)
    return result, nil
}

func (this *Server) CreateProject(ctx context.Context, req *pb.Project) (*pb.Project, error) {
	log.Info("got request for ID:", req.Id)
	project, err := this.projects.CreateProject(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Failed creating project: %w", err)
	}
    log.Info("project was created: ", project)
	return project, nil
}

func (this *Server) serve(addr string) (error) {
	this.projects = *NewProjects()
	log.Info("Listening on ", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	var opts []grpc.ServerOption

	server := grpc.NewServer(opts...)
	log.Info("Registering service")
	pb.RegisterProjectServiceServer(server, this)
	log.Info("Serving...")
	server.Serve(lis)
	return nil
}