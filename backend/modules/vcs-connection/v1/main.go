package vcs_connection

import (
	"alfred/config"
	"alfred/modules/vcs-connection/v1/internals"
	internal_pb "alfred/modules/vcs-connection/v1/internals/pb"
	"alfred/modules/vcs-connection/v1/models"
	"alfred/modules/vcs-connection/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type vcsConnectionServer struct {
	db                *gorm.DB
	config            *config.Config
	grpcClient        *grpc.ClientConn
	internalVCSClient internal_pb.VCSConnectionInternalServer
}

// NewVCSConnectionServer todo : AlWays add migration code for best practices
func NewVCSConnectionServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.VCSConnectionsServer {

	//initial migration of databases: schema migration
	models.InitialMigrationVCSConnection(db)
	internalVCSClient := internals.NewVCSConnectionInternalServer(db, config, grpcClientConn)
	return &vcsConnectionServer{
		db:                db,
		config:            config,
		grpcClient:        grpcClientConn,
		internalVCSClient: internalVCSClient,
	}
}
