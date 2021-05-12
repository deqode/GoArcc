package internals

import (
	"alfred/config"
	accountPb "alfred/modules/account/v1/pb"
	"alfred/modules/user-profile/v1/internals/pb"
	"alfred/modules/user-profile/v1/models"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type userProfileInternalServer struct {
	db            *gorm.DB
	config        *config.Config
	grpcClient    *grpc.ClientConn
	accountClient accountPb.AccountsServer
}

// NewUserProfileInternalServer todo : AlWays add migration code for best practices
func NewUserProfileInternalServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.UserProfileInternalServer {

	//initial migration of databases: schema migration
	models.InitialMigrationUserProfile(db)
	return &userProfileInternalServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
