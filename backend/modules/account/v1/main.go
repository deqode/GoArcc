package account

import (
	"alfred/config"
	"alfred/modules/account/v1/models"
	"alfred/modules/account/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type accountsServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

// NewAccountsServer todo : AlWays add migration code for best practices
func NewAccountsServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.AccountsServer {

	//initial migration of databases: schema migration
	models.InitialMigrationAccount(db)
	return &accountsServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
