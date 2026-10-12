package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	go_load "go_load/internal/generated"
)

type Handler struct {
	go_load.UnimplementedGoLoadServiceServer
}

func NewHandler() go_load.GoLoadServiceServer {
	return &Handler{}
}

func (h *Handler) CreateAccount(_ context.Context, _ *go_load.CreateAccountRequest) (*go_load.CreateAccountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateAccount not implemented")
}

func (h *Handler) CreateSession(_ context.Context, _ *go_load.CreateSessionRequest) (*go_load.CreateSessionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateSession not implemented")
}

func (h *Handler) CreateDownloadTask(_ context.Context, _ *go_load.CreateDownloadTaskRequest) (*go_load.CreateDownloadTaskResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateDownloadTask not implemented")
}

func (h *Handler) GetDownloadTaskList(_ context.Context, _ *go_load.GetDownloadTaskListRequest) (*go_load.GetDownloadTaskListResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetDownloadTaskList not implemented")
}

func (h *Handler) UpdateDownloadTask(_ context.Context, _ *go_load.UpdateDownloadTaskRequest) (*go_load.UpdateDownloadTaskResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method UpdateDownloadTask not implemented")
}

func (h *Handler) DeleteDownloadTask(_ context.Context, _ *go_load.DeleteDownloadTaskRequest) (*go_load.DeleteDownloadTaskResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method DeleteDownloadTask not implemented")
}

func (h *Handler) GetDownloadTaskFile(_ *go_load.GetDownloadTaskFileRequest, _ grpc.ServerStreamingServer[go_load.GetDownloadTaskFileResponse]) error {
	return status.Error(codes.Unimplemented, "method GetDownloadTaskFile not implemented")
}
