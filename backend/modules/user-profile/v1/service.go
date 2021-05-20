package user_profile

import (
	"alfred/config"
	pb2 "alfred/modules/account/v1/external-svc/pb"
	model "alfred/modules/user-profile/v1/internal/models"
	"alfred/modules/user-profile/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type userProfilesServer struct {
	db            *gorm.DB
	config        *config.Config
	grpcClient    *grpc.ClientConn
	accountClient pb2.AccountsServer
	store         model.UserProfileStore
}

// NewUserProfilesServer todo : AlWays add migration code for best practices
func NewUserProfilesServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.UserProfilesServer {

	//initial migration of databases: schema migration
	model.InitialMigrationUserProfile(db)
	UserProfileStore := model.NewUserProfileStore(db)
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
	accountClient pb2.AccountsServer
	store         model.UserProfileStore
}

func NewUserProfileInServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.UserProfileInternalServer {

	UserProfileStore := model.NewUserProfileStore(db)
	return &userProfileInServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
		store:      UserProfileStore,
	}
}
