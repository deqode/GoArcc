package VCSConnectionService

import (
	"alfred/background/workers/workflows/vcs/git"
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

	err := s.triggerGitClone(git.GitCloneRequest{
		AccountId:     "",
		UserId:        "",
		RepositoryURL: "https://github.com/nndd91/cadence-api-example.git",
		BranchName:    "",
	})
	if err != nil {
		return nil, err
	}

	return &pb.ListAllSupportedVCSProvidersResponse{
		Providers: providers,
	}, nil
}
func (s *VCSConnectionService) ListVCSConnection(context.Context, *pb.ListVCSConnectionRequest) (*pb.ListVCSConnectionResponse, error) {
	return nil, nil
}
