package external_svc

import (
	"alfred/config"
	accountpb "alfred/modules/account/v1/pb"
	"alfred/modules/user-profile/v1/internal-svc"
	"alfred/modules/user-profile/v1/models"
	"alfred/modules/user-profile/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type userProfilesServer struct {
	db            *gorm.DB
	config        *config.Config
	grpcClient    *grpc.ClientConn
	accountClient accountpb.AccountsServer
	store         internal_svc.UserProfileStore
}

// NewUserProfilesServer todo : AlWays add migration code for best practices
func NewUserProfilesServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.UserProfilesServer {

	//initial migration of databases: schema migration
	internal_svc.InitialMigrationUserProfile(db)
	UserProfileStore := models.NewUserProfileStore(db)
	return &userProfilesServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
		store:      UserProfileStore,
	}
}

// Internal Service Configuration
type userProfileInServer struct {
	db            *gorm.DB
	config        *config.Config
	grpcClient    *grpc.ClientConn
	accountClient accountpb.AccountsServer
	store         internal_svc.UserProfileStore
}

func NewUserProfileInServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.UserProfileInternalServer {

	UserProfileStore := models.NewUserProfileStore(db)
	return &userProfileInServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
		store:      UserProfileStore,
	}
}
