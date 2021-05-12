package account

import (
	"alfred/config"
	"alfred/modules/account/v1/internals/pb"
	"alfred/modules/account/v1/models"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type accountInternalServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

// NewAccountInternalServer todo : AlWays add migration code for best practices
func NewAccountInternalServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.AccountInternalServer {

	//initial migration of databases: schema migration
	models.InitialMigrationAccount(db)
	return &accountInternalServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
