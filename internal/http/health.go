package http

import (
	"context"

	piper "github.com/dreamstax/go-genproto/kaiapis/piper/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) HealthCheck(ctx context.Context, req *piper.HealthCheckRequest) (*piper.HealthCheckResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not yet implemented")
}
