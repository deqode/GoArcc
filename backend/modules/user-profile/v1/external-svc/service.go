package external_svc

import (
	"alfred/config"
	accountPb "alfred/modules/account/v1/pb"
	model "alfred/modules/user-profile/v1/models"
	"alfred/modules/user-profile/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type userProfilesServer struct {
	db            *gorm.DB
	config        *config.Config
	grpcClient    *grpc.ClientConn
	accountServer accountPb.AccountsServer
}

// NewUserProfilesServer todo : AlWays add migration code for best practices
func NewUserProfilesServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.UserProfilesServer {

	//initial migration of databases: schema migration
	model.InitialMigrationUserProfile(db)
	return &userProfilesServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
