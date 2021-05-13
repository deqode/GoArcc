package git

import (
	githubService "alfred/modules/git/v1/github"
	"alfred/modules/git/v1/pb"
	vcsinternalPb "alfred/modules/vcs-connection/v1/internals/pb"
	"alfred/protos/types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//GetRepository get a repository
func (s *gitServer) GetRepository(ctx context.Context, in *pb.GetRepositoryRequest) (*pb.Repository, error) {
	vcs, err := s.vcsInServer.GetVCSConnection(ctx, &vcsinternalPb.GetVCSConnectionRequest{
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
		return nil, status.Error(codes.NotFound, "Access token not found")
	}
	switch in.Provider {
	case types.VCSProviders_GITHUB:
		client := githubService.NewGithubService(ctx, accessToken)
		return client.GetRepository(ctx, in)
	}
	return nil, nil
}
