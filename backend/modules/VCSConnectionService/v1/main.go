package v1

import (
	"alfred/config"
	"alfred/modules/VCSConnectionService/v1/models"
	"alfred/modules/VCSConnectionService/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"path/filepath"
)

type VCSConnectionService struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
	info       []*Info
}

//todo : AlWays add migration code for best practices
func NewVCSConnectionService(db *gorm.DB, config *config.Config, grpcClientConn *grpc.ClientConn) pb.VCSConnectionServiceServer {

	//initial migration of databases: schema migration
	models.InitialMigrationVCSConnection(db)
	abs, _ := filepath.Abs("./modules/VCSConnectionService/info.json")
	vcsConfig := InfoFromFile(abs)
	return &VCSConnectionService{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
		info:       vcsConfig,
	}
}
