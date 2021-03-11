package UserProfileService

import (
	"alfred/config"
	"alfred/modules/UserProfileService/models"
	"alfred/modules/UserProfileService/pb"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type UserProfileService struct {
	fx.In
	db     *gorm.DB
	config *config.Config
	/*	helloCli hellopb.HelloWorldServiceClient*/
}

//todo : AlWays add migration code for best practices
func NewUserProfileService(db *gorm.DB /* helloCli hellopb.HelloWorldServiceClient*/, config *config.Config) pb.UserProfileServiceServer {

	//initial migration of databases: schema migration
	models.InitialMigrationUserProfile(db)
	return &UserProfileService{
		db:     db,
		config: config,
	}
}
