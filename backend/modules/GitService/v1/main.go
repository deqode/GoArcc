package GitService

import (
	"alfred/config"
	"alfred/modules/GitService/v1/models"
	"alfred/modules/GitService/v1/pb"
	"github.com/google/go-github/v34/github"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type GitService struct {
	db           *gorm.DB
	config       *config.Config
	grpcClient   *grpc.ClientConn
	githubClient *github.Client
}

//todo : AlWays add migration code for best practices
func NewGitService(db *gorm.DB, config *config.Config, grpcClientConn *grpc.ClientConn) pb.GitServiceServer {

	//initial migration of databases: schema migration
	github.NewClient()
	return &GitService{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
