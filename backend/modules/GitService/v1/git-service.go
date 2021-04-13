package GitService

import (
	githubService "alfred/modules/GitService/v1/github"
	"alfred/modules/GitService/v1/pb"
	vcsinternalPb "alfred/modules/VCSConnectionService/v1/internals/pb"
	"alfred/protos/types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GitService) ListReposistory(ctx context.Context, in *pb.ListReposistoryRequest) (*pb.ListReposistoryResponse, error) {
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
		githubService.NewGithubService(ctx, accessToken)
		return s.githubService.ListReposistory(ctx, in)
	}
	return nil, nil
}
func (s *GitService) GetReposistory(ctx context.Context, in *pb.GetReposistoryRequest) (*pb.Reposistory, error) {
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
		githubService.NewGithubService(ctx, accessToken)
		return s.githubService.GetReposistory(ctx, in)
	}
	return nil, nil
}
