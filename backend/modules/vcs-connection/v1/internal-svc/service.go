package internal_svc

import (
	"alfred/config"
	"alfred/modules/vcs-connection/v1/models"
	"alfred/modules/vcs-connection/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type vcsConnectionIntServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

// NewVCSConnectionIntServer todo : AlWays add migration code for best practices
func NewVCSConnectionIntServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.VCSConnectionInternalServer {
	//initial migration of databases: schema migration
	models.InitialMigrationVCSConnection(db)
	return &vcsConnectionIntServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
