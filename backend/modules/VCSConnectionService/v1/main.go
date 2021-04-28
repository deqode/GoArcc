package VCSConnectionService

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	"alfred/modules/VCSConnectionService/v1/internals"
	internal_pb "alfred/modules/VCSConnectionService/v1/internals/pb"
	"alfred/modules/VCSConnectionService/v1/models"
	"alfred/modules/VCSConnectionService/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type VCSConnectionService struct {
	db                *gorm.DB
	config            *config.Config
	grpcClient        *grpc.ClientConn
	internalVCSClient internal_pb.VCSConnectionInternalServiceServer
	cadenceConfig     *background.CadenceAppConfig
	cadenceAdapter    *cadenceAdapter.CadenceAdapter
}

// NewVCSConnectionService todo : AlWays add migration code for best practices
func NewVCSConnectionService(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
	cadenceConfig *background.CadenceAppConfig,
	cadenceAdapter *cadenceAdapter.CadenceAdapter) pb.VCSConnectionServiceServer {

	//initial migration of databases: schema migration
	models.InitialMigrationVCSConnection(db)
	internalVCSClient := internals.NewVCSConnectionInternalService(db, config, grpcClientConn)
	return &VCSConnectionService{
		db:                db,
		config:            config,
		grpcClient:        grpcClientConn,
		internalVCSClient: internalVCSClient,
		cadenceConfig:     cadenceConfig,
		cadenceAdapter:    cadenceAdapter,
	}
}
