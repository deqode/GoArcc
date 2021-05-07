package GitService

import (
	githubService "alfred/modules/GitService/v1/github"
	"alfred/modules/GitService/v1/pb"
	vcsinternalPb "alfred/modules/VCSConnection/v1/internals/pb"
	"alfred/protos/types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//ListRepository List All the repository
func (s *GitService) ListRepository(ctx context.Context, in *pb.ListRepositoryRequest) (*pb.ListRepositoryResponse, error) {
	in.Provider = types.GitProviders_GITHUB
	//fetch the stored accessToken
	vcs, err := s.vcsInternalServer.GetVCSConnection(ctx, &vcsinternalPb.GetVCSConnectionRequest{
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
	case types.GitProviders_GITHUB:
		client := githubService.NewGithubService(ctx, accessToken)
		return client.ListRepository(ctx, in)
		//return githubService.ListRepository(ctx, in)
	}
	return nil, nil
}

func (s *GitService) GetRepository(ctx context.Context, in *pb.GetRepositoryRequest) (*pb.Repository, error) {
	//in.Provider = types.GitProviders_GITHUB
	//in.RepoName = "chatbox"
	vcs, err := s.vcsInternalServer.GetVCSConnection(ctx, &vcsinternalPb.GetVCSConnectionRequest{
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
	case types.GitProviders_GITHUB:
		client := githubService.NewGithubService(ctx, accessToken)
		return client.GetRepository(ctx, in)
	}
	return nil, nil
}
