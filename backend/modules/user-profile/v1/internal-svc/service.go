package internal_svc

import (
	"alfred/config"
	accountPb "alfred/modules/account/v1/pb"
	model "alfred/modules/user-profile/v1/models"
	"alfred/modules/user-profile/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// Internal Service Configuration
type userProfileInServer struct {
	db            *gorm.DB
	config        *config.Config
	grpcClient    *grpc.ClientConn
	accountClient accountPb.AccountsServer
}

func NewUserProfileInServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.UserProfileInternalServer {
	//initial migration of databases: schema migration
	model.InitialMigrationUserProfile(db)
	return &userProfileInServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
