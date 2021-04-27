package UserProfileService

import (
	"alfred/config"
	"alfred/modules/AccountService/v1"
	accountPb "alfred/modules/AccountService/v1/pb"
	"alfred/modules/UserProfileService/v1/models"
	"alfred/modules/UserProfileService/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type UserProfileService struct {
	db            *gorm.DB
	config        *config.Config
	grpcClient    *grpc.ClientConn
	accountClient accountPb.AccountServiceServer
}

//todo : AlWays add migration code for best practices
func NewUserProfileService(db *gorm.DB, config *config.Config, grpcClientConn *grpc.ClientConn) pb.UserProfileServiceServer {

	//initial migration of databases: schema migration
	models.InitialMigrationUserProfile(db)
	accountCli := AccountService.NewAccountService(db, config, grpcClientConn)
	return &UserProfileService{
		db:            db,
		config:        config,
		grpcClient:    grpcClientConn,
		accountClient: accountCli,
	}
}
