package VCSConnectionService

import (
	"alfred/config"
	"alfred/modules/VCSConnectionService/models"
	"alfred/modules/VCSConnectionService/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type VCSConnectionService struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

//todo : AlWays add migration code for best practices
func NewVCSConnectionService(db *gorm.DB, config *config.Config, grpcClientConn *grpc.ClientConn) pb.VCSConnectionServiceServer {

	//initial migration of databases: schema migration
	models.InitialMigrationVCSConnection(db)
	return &VCSConnectionService{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
