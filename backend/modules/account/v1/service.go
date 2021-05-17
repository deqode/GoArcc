package account

import (
	"alfred/config"
	model "alfred/modules/account/v1/internal/models"
	"alfred/modules/account/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type accountsServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
	store      model.AccountStore
}

// NewAccountsServer todo : AlWays add migration code for best practices
func NewAccountsServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.AccountsServer {

	//initial migration of databases: schema migration
	model.InitialMigrationAccount(db)
	accountStore := model.NewAccountStore(db)
	return &accountsServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
		store:      accountStore,
	}
}
