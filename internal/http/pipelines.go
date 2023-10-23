package http

import (
	"context"

	piper "github.com/dreamstax/go-genproto/kaiapis/piper/v1alpha1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListPipelines(ctx context.Context, req *piper.ListPipelinesRequest) (*piper.ListPipelinesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not yet implemented")
}

func (s *Server) GetPipeline(ctx context.Context, req *piper.GetPipelineRequest) (*piper.Pipeline, error) {
	return nil, status.Error(codes.Unimplemented, "not yet implemented")
}

func (s *Server) CreatePipeline(ctx context.Context, req *piper.CreatePipelineRequest) (*piper.Pipeline, error) {
	return nil, status.Error(codes.Unimplemented, "not yet implemented")
}

func (s *Server) UpdatePipeline(ctx context.Context, req *piper.UpdatePipelineRequest) (*piper.Pipeline, error) {
	return nil, status.Error(codes.Unimplemented, "not yet implemented")
}

func (s *Server) DeletePipeline(ctx context.Context, req *piper.DeletePipelineRequest) (*empty.Empty, error) {
	return &emptypb.Empty{}, status.Error(codes.Unimplemented, "not yet implemented")
}
