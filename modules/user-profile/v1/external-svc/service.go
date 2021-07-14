package external_svc

import (
	"github.com/deqode/GoArcc/config"
	model "github.com/deqode/GoArcc/modules/user-profile/v1/models"
	"github.com/deqode/GoArcc/modules/user-profile/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type userProfilesServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

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
