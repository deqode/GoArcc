package vcs_connection

import (
	"alfred/config"
	model "alfred/modules/vcs-connection/v1/internal/models"
	"alfred/modules/vcs-connection/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type vcsConnectionServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

// NewVCSConnectionServer todo : AlWays add migration code for best practices
func NewVCSConnectionServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.VCSConnectionsServer {
	//initial migration of databases: schema migration
	model.InitialMigrationVCSConnection(db)
	return &vcsConnectionServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
