package external_svc

import (
	"alfred/config"
	"alfred/modules/account/v1/controllers"
	extPb "alfred/modules/account/v1/external-svc/pb"
	"alfred/modules/account/v1/models"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type accountExtServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
	store      controllers.AccountStore
}

// NewAccountExtServer todo : AlWays add migration code for best practices
func NewAccountExtServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) extPb.AccountsServer {

	//initial migration of databases: schema migration
	models.InitialMigrationAccount(db)
	accountStore := controllers.NewAccountStore(db)
	return &accountExtServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
		store:      accountStore,
	}
}
