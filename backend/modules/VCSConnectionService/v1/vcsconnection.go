package VCSConnectionService

import (
	internal_pb "alfred/modules/VCSConnectionService/v1/internals/pb"
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
func (s *VCSConnectionService) ListVCSConnection(ctx context.Context, in *pb.ListVCSConnectionRequest) (*pb.ListVCSConnectionResponse, error) {
	vcs, err := s.internalVCSClient.GetVCSConnection(ctx, &internal_pb.GetVCSConnectionRequest{
		AccountId: in.AccountId,
		Provider:  in.Provider,
	})
	if err != nil {
		return nil, err
	}
	accountVcsConnection := &pb.AccountVCSConnection{
		Provider:  vcs.Provider,
		AccountId: vcs.AccountId,
		Label:     vcs.Label,
		UserName:  vcs.UserName,
	}
	return &pb.ListVCSConnectionResponse{
		VcsConnection: []*pb.AccountVCSConnection{accountVcsConnection},
	}, nil
}
