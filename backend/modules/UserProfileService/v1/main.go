package v1

import (
	"alfred/config"
	"alfred/modules/UserProfileService/v1/models"
	"alfred/modules/UserProfileService/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type UserProfileService struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

//todo : AlWays add migration code for best practices
func NewUserProfileService(db *gorm.DB, config *config.Config, grpcClientConn *grpc.ClientConn) pb.UserProfileServiceServer {

	//initial migration of databases: schema migration
	models.InitialMigrationUserProfile(db)
	return &UserProfileService{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
