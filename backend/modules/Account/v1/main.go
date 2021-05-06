package Account

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	"alfred/modules/Account/v1/models"
	"alfred/modules/Account/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type accountsServer struct {
	db             *gorm.DB
	config         *config.Config
	grpcClient     *grpc.ClientConn
	cadenceConfig  *background.CadenceAppConfig
	cadenceAdapter *cadenceAdapter.CadenceAdapter
}

// NewAccountsServer todo : AlWays add migration code for best practices
func NewAccountsServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
	cadenceConfig *background.CadenceAppConfig,
	cadenceAdapter *cadenceAdapter.CadenceAdapter,
) pb.AccountsServer {

	//initial migration of databases: schema migration
	models.InitialMigrationAccount(db)
	return &accountsServer{
		db:             db,
		config:         config,
		grpcClient:     grpcClientConn,
		cadenceAdapter: cadenceAdapter,
		cadenceConfig:  cadenceConfig,
	}
}
