package github

import (
	"alfred/modules/git/v1/pb"
	"context"
	"github.com/google/go-github/v34/github"
	"golang.org/x/oauth2"
)

type Service struct {
	githubClient *github.Client
}

func NewGithubService(ctx context.Context, accessToken string) pb.GitsServer {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return &Service{
		githubClient: client,
	}
}

func (s *Service) ListRepository(ctx context.Context, in *pb.ListRepositoryRequest) (*pb.ListRepositoryResponse, error) {
	var repositories []*pb.Repository
	repos, _, err := s.githubClient.Repositories.List(ctx, "", nil)
	if err != nil {
		return nil, err
	}
	for _, repo := range repos {
		if !*repo.Private {
			Repository := &pb.Repository{
				Id:            *repo.ID,
				NodeId:        *repo.NodeID,
				Name:          *repo.Name,
				FullName:      *repo.FullName,
				DefaultBranch: *repo.DefaultBranch,
				CloneUrl:      *repo.CloneURL,
				GitUrl:        *repo.GitURL,
				Private:       *repo.Private,
				Branches:      nil,
			}
			repositories = append(repositories, Repository)
		}
	}
	return &pb.ListRepositoryResponse{
		Repositories: repositories,
	}, nil
}

func (s *Service) GetRepository(ctx context.Context, in *pb.GetRepositoryRequest) (*pb.Repository, error) {
	//get Repository with branches
	repo, _, err := s.githubClient.Repositories.Get(ctx, in.OwnerName, in.RepoName)
	if err != nil {
		return nil, err
	}

	//get Repository branches
	res, _, err := s.githubClient.Repositories.ListBranches(ctx, in.OwnerName, in.RepoName, nil)
	if err != nil {
		return nil, err
	}
	var branches []string
	for _, br := range res {
		branches = append(branches, *br.Name)
	}
	Repository := &pb.Repository{
		Id:            *repo.ID,
		NodeId:        *repo.NodeID,
		Name:          *repo.Name,
		FullName:      *repo.FullName,
		DefaultBranch: *repo.DefaultBranch,
		CloneUrl:      *repo.CloneURL,
		GitUrl:        *repo.GitURL,
		Private:       *repo.Private,
		Branches:      branches,
	}
	return Repository, nil
}

func (s *Service) CloneRepository(ctx context.Context, in *pb.CloneRepositoryRequest) (*pb.CloneRepositoryResponse, error) {
	panic("implement me")
}

func (s *Service) GetCloningStatus(ctx context.Context, in *pb.GetCloningStatusRequest) (*pb.GetCloningStatusResponse, error) {
	panic("implement me")
}
