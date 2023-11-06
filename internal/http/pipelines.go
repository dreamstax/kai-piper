package http

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	piper "github.com/dreamstax/go-genproto/dreamstaxapis/piper/v1alpha1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListPipelineJobs(ctx context.Context, req *piper.ListPipelineJobsRequest) (*piper.ListPipelineJobsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not yet implemented")
}

func (s *Server) GetPipelineJob(ctx context.Context, req *piper.GetPipelineJobRequest) (*piper.PipelineJob, error) {
	return nil, status.Error(codes.Unimplemented, "not yet implemented")
}

func (s *Server) CreatePipelineJob(ctx context.Context, req *piper.CreatePipelineJobRequest) (*piper.PipelineJob, error) {
	return nil, status.Error(codes.Unimplemented, "not yet implemented")
}

func (s *Server) UpdatePipelineJob(ctx context.Context, req *piper.UpdatePipelineJobRequest) (*piper.PipelineJob, error) {
	return nil, status.Error(codes.Unimplemented, "not yet implemented")
}

func (s *Server) DeletePipelineJob(ctx context.Context, req *piper.DeletePipelineJobRequest) (*empty.Empty, error) {
	return &emptypb.Empty{}, status.Error(codes.Unimplemented, "not yet implemented")
}

func (s *Server) RunPipelineJob(ctx context.Context, req *piper.RunPipelineJobRequest) (*longrunningpb.Operation, error) {
	return nil, status.Error(codes.Unimplemented, "not yet implemented")
}

func (s *Server) ListPipelineJobExecutions(ctx context.Context, req *piper.ListPipelineJobExecutionsRequest) (*piper.ListPipelineJobExecutionsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not yet implemented")
}

func (s *Server) GetPipelineJobExecution(ctx context.Context, req *piper.GetPipelineJobExecutionRequest) (*piper.PipelineJobExecution, error) {
	return nil, status.Error(codes.Unimplemented, "not yet implemented")
}

func (s *Server) DeletePipelineJobExecution(ctx context.Context, req *piper.DeletePipelineJobExecutionRequest) (*empty.Empty, error) {
	return &emptypb.Empty{}, status.Error(codes.Unimplemented, "not yet implemented")
}
