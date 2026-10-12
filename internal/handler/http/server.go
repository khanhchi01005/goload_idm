package http

import (
	"context"
	go_load "go_load/internal/generated"
	"log"
	"net"
	"net/http"
	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
type Server interface{
	Start(ctx context.Context) error
}

type server struct{

}

func NewServer() Server{
	return &server{

	}
}

func (s *server) Start(ctx context.Context) error{
	mux := runtime.NewServeMux()
	err := go_load.RegisterGoLoadServiceHandlerFromEndpoint(
		ctx,
		mux,
		"0.0.0.0:8000",
		[]grpc.DialOption(insecure.NewCredentials()),
	); err != nil {
		return err
	}

	return http.ListenAndServe(":8001", mux)
}