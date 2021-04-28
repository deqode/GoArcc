package AccountService

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	"alfred/modules/AccountService/v1/models"
	"alfred/modules/AccountService/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type AccountService struct {
	db             *gorm.DB
	config         *config.Config
	grpcClient     *grpc.ClientConn
	cadenceConfig  *background.CadenceAppConfig
	cadenceAdapter *cadenceAdapter.CadenceAdapter
}

// NewAccountService todo : AlWays add migration code for best practices
func NewAccountService(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
	cadenceConfig *background.CadenceAppConfig,
	cadenceAdapter *cadenceAdapter.CadenceAdapter,
) pb.AccountServiceServer {

	//initial migration of databases: schema migration
	models.InitialMigrationAccount(db)
	return &AccountService{
		db:             db,
		config:         config,
		grpcClient:     grpcClientConn,
		cadenceAdapter: cadenceAdapter,
		cadenceConfig:  cadenceConfig,
	}
}
