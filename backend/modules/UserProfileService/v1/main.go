package UserProfileService

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	"alfred/modules/Account/v1"
	accountPb "alfred/modules/Account/v1/pb"
	"alfred/modules/UserProfileService/v1/models"
	"alfred/modules/UserProfileService/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type UserProfileService struct {
	db             *gorm.DB
	config         *config.Config
	grpcClient     *grpc.ClientConn
	accountClient  accountPb.AccountsServer
	cadenceConfig  *background.CadenceAppConfig
	cadenceAdapter *cadenceAdapter.CadenceAdapter
}

// NewUserProfileService todo : AlWays add migration code for best practices
func NewUserProfileService(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
	cadenceConfig *background.CadenceAppConfig,
	cadenceAdapter *cadenceAdapter.CadenceAdapter,
) pb.UserProfileServiceServer {

	//initial migration of databases: schema migration
	models.InitialMigrationUserProfile(db)
	accountCli := Account.NewAccountsServer(db, config, grpcClientConn, cadenceConfig, cadenceAdapter)
	return &UserProfileService{
		db:             db,
		config:         config,
		grpcClient:     grpcClientConn,
		accountClient:  accountCli,
		cadenceAdapter: cadenceAdapter,
		cadenceConfig:  cadenceConfig,
	}
}
