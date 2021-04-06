package v1

import (
	"alfred/config"
	internal "alfred/modules/VCSConnectionService/v1/internal/pb"
	"alfred/modules/VCSConnectionService/v1/models"
	"alfred/modules/VCSConnectionService/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type VCSConnectionService struct {
	db                *gorm.DB
	config            *config.Config
	grpcClient        *grpc.ClientConn
	internalVCSClient internal.VCSConnectionInternalServiceClient
}

//todo : AlWays add migration code for best practices
func NewVCSConnectionService(db *gorm.DB, config *config.Config, grpcClientConn *grpc.ClientConn) pb.VCSConnectionServiceServer {

	//initial migration of databases: schema migration
	models.InitialMigrationVCSConnection(db)
	internalVCSClient := internal.NewVCSConnectionInternalServiceClient(grpcClientConn)
	return &VCSConnectionService{
		db:                db,
		config:            config,
		grpcClient:        grpcClientConn,
		internalVCSClient: internalVCSClient,
	}
}
