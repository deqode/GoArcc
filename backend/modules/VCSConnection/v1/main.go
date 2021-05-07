package vcs_connection

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	"alfred/modules/VCSConnection/v1/internals"
	internal_pb "alfred/modules/VCSConnection/v1/internals/pb"
	"alfred/modules/VCSConnection/v1/models"
	"alfred/modules/VCSConnection/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type vcsConnectionServer struct {
	db                *gorm.DB
	config            *config.Config
	grpcClient        *grpc.ClientConn
	internalVCSClient internal_pb.VCSConnectionInternalServer
	cadenceConfig     *background.CadenceAppConfig
	cadenceAdapter    *cadenceAdapter.CadenceAdapter
}

// NewVCSConnectionServer todo : AlWays add migration code for best practices
func NewVCSConnectionServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
	cadenceConfig *background.CadenceAppConfig,
	cadenceAdapter *cadenceAdapter.CadenceAdapter) pb.VCSConnectionsServer {

	//initial migration of databases: schema migration
	models.InitialMigrationVCSConnection(db)
	internalVCSClient := internals.NewVCSConnectionInternalServer(db, config, grpcClientConn)
	return &vcsConnectionServer{
		db:                db,
		config:            config,
		grpcClient:        grpcClientConn,
		internalVCSClient: internalVCSClient,
		cadenceConfig:     cadenceConfig,
		cadenceAdapter:    cadenceAdapter,
	}
}
