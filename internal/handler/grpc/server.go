package grpc

import (
	"context"
	go_load "go_load/internal/generated"
	"log"
	"net"

	"google.golang.org/grpc"
)


type Server interface {
	Start(ctx context.Context) error
}

type server struct {
	handler go_load.GoLoadServiceServer
}


func NewServer(handler go_load.GoLoadServiceServer) Server {
	return &server{
		handler: handler,
	} 
}

func (s *server) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatalf("failed to open tcp port: %v", err)
	}

	grpcServer := grpc.NewServer()

	go_load.RegisterGoLoadServiceServer(grpcServer, s.handler)
	return grpcServer.Serve(lis)
}