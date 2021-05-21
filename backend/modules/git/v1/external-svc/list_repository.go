package external_svc

import (
	github2 "alfred/modules/git/v1/internal-svc/github"
	"alfred/modules/git/v1/pb"
	vcsPb "alfred/modules/vcs-connection/v1/pb"
	"alfred/protos/types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//ListRepository List All the repository
func (s *gitServer) ListRepository(ctx context.Context, in *pb.ListRepositoryRequest) (*pb.ListRepositoryResponse, error) {
	//Validate request
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	vcs, err := s.vcsInServer.GetVCSConnection(ctx, &vcsPb.GetVCSConnectionRequest{
		AccountId: in.AccountId,
		Provider:  in.Provider,
	})
	if err != nil {
		return nil, err
	}
	var accessToken string
	if vcs.AccessToken != "" {
		accessToken = vcs.AccessToken
	} else {
		return nil, status.Error(codes.NotFound, "Integration token not found")
	}
	switch in.Provider {
	case types.VCSProviders_GITHUB:
		client := github2.NewGithubService(ctx, accessToken)
		return client.ListRepository(ctx, in)
	}
	return nil, nil
}
