package internal_svc

import (
	"alfred/config"
	"alfred/modules/account/v1/models"
	intPb "alfred/modules/account/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

//Account Internal Server Configuration
type accountsIntServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

func NewAccountsInServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) intPb.AccountInternalServer {

	//initial migration of databases: schema migration
	models.InitialMigrationAccount(db)
	return &accountsIntServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
