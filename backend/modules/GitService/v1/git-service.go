package GitService

import (
	"alfred/modules/GitService/v1/pb"
	vcsinternalPb "alfred/modules/VCSConnectionService/v1/internal/pb"
	"context"
	"github.com/google/go-github/v34/github"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GitService) ListReposistory(ctx context.Context, in *pb.ListReposistoryRequest) (*pb.ListReposistoryResponse, error) {
	//check the stored accessToken
	vcs, err := s.vcsConnectionClient.GetVCSConnection(ctx, &vcsinternalPb.GetVCSConnectionRequest{
		AccountId: "",
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

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	var reposistories []*pb.Reposistory
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		return nil, err
	}
	for _, repo := range repos {
		reposistory := &pb.Reposistory{
			Id:       *repo.ID,
			NodeID:   *repo.NodeID,
			Name:     *repo.Name,
			FullName: *repo.FullName,
			//Description:   *repo.Description,
			DefaultBranch: *repo.DefaultBranch,
			//MasterBranch:  *repo.MasterBranch,
			//CreatedAt:     repo.CreatedAt,
			//PushedAt:      *repo.PushedAt,
			//UpdatedAt:     *repo.UpdatedAt,
			Clone_URL: *repo.CloneURL,
			Git_URL:   *repo.GitURL,
			//Size:          *repo.Size,
			Private:  *repo.Private,
			Branches: nil,
		}
		reposistories = append(reposistories, reposistory)
	}
	return &pb.ListReposistoryResponse{
		Reposistories: reposistories,
	}, nil
}
func (s *GitService) GetReposistory(ctx context.Context, in *pb.GetReposistoryRequest) (*pb.Reposistory, error) {
	return nil, nil
}
