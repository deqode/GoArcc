package user_profile

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	"alfred/modules/account/v1"
	accountPb "alfred/modules/account/v1/pb"
	"alfred/modules/user-profile/v1/models"
	"alfred/modules/user-profile/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type userProfileServer struct {
	db             *gorm.DB
	config         *config.Config
	grpcClient     *grpc.ClientConn
	accountClient  accountPb.AccountsServer
	cadenceConfig  *background.CadenceAppConfig
	cadenceAdapter *cadenceAdapter.CadenceAdapter
}

// NewUserProfileServer todo : AlWays add migration code for best practices
func NewUserProfileServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
	cadenceConfig *background.CadenceAppConfig,
	cadenceAdapter *cadenceAdapter.CadenceAdapter,
) pb.UserProfileServiceServer {

	//initial migration of databases: schema migration
	models.InitialMigrationUserProfile(db)
	accountCli := account.NewAccountsServer(db, config, grpcClientConn, cadenceConfig, cadenceAdapter)
	return &userProfileServer{
		db:             db,
		config:         config,
		grpcClient:     grpcClientConn,
		accountClient:  accountCli,
		cadenceAdapter: cadenceAdapter,
		cadenceConfig:  cadenceConfig,
	}
}
