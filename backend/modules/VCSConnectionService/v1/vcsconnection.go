package VCSConnectionService

import (
	"alfred/modules/VCSConnectionService/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

var (
	GithubOAuthScope = []string{"user", "gist", "repo", "repo_deployment", "repo_hook"}
)

func (s *VCSConnectionService) ListAllSupportedVCSProviders(context.Context, *empty.Empty) (*pb.ListAllSupportedVCSProvidersResponse, error) {
	conf := s.config.VCSConfig
	var providers []string
	for _, vcs := range conf {
		providers = append(providers, vcs.Provider)
	}
	return &pb.ListAllSupportedVCSProvidersResponse{
		Providers: providers,
	}, nil
}
func (s *VCSConnectionService) ListVCSConnection(context.Context, *pb.ListVCSConnectionRequest) (*pb.ListVCSConnectionResponse, error) {
	return nil, nil
}
