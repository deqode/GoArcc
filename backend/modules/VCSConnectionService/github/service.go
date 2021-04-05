package github

import (
	"alfred/modules/VCSConnectionService/pb"
	"context"
	"github.com/google/go-github/v34/github"
	"golang.org/x/oauth2"
)

type Service struct {
	githubClient *github.Client
}

func NewGithubService(ctx context.Context, accessToken string) Service {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return Service{
		githubClient: client,
	}
}
func (s *Service) ListReposistory(ctx context.Context, in *pb.ListReposistoryRequest) (*pb.ListReposistoryResponse, error) {
	var reposistories []*pb.Reposistory
	repos, _, err := s.githubClient.Repositories.List(ctx, "", nil)
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

func (s *Service) GetReposistory(ctx context.Context, in *pb.GetReposistoryRequest) (*pb.Reposistory, error) {
	//get reposistory with branches
	repo, _, err := s.githubClient.Repositories.Get(ctx, "", in.RepoName)
	if err != nil {
		return nil, err
	}

	res, _, err := s.githubClient.Repositories.ListBranches(ctx, "", in.RepoName, nil)
	if err != nil {
		return nil, err
	}
	var branches []string
	for _, br := range res {
		branches = append(branches, *br.Name)
	}
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
		Branches: branches,
	}
	return reposistory, nil
}
