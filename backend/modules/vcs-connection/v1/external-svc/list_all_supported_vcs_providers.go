package external_svc

import (
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *vcsConnectionServer) ListAllSupportedVCSProviders(context.Context, *empty.Empty) (*pb.ListAllSupportedVCSProvidersResponse, error) {
	conf := s.config.SupportedVcsConfig
	if conf == nil {
		return nil, status.Error(codes.FailedPrecondition, "No VCS Providers")
	}
	return &pb.ListAllSupportedVCSProvidersResponse{
		Providers: conf,
	}, nil
}
