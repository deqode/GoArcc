package internal_svc

import (
	"github.com/deqode/GoArcc/config"
	model "github.com/deqode/GoArcc/modules/user-profile/v1/models"
	"github.com/deqode/GoArcc/modules/user-profile/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// Internal Service Configuration
type userProfileInServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
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
