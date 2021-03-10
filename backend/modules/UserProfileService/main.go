package UserProfileService

import (
	hellopb "alfred/modules/HelloWorldService/pb"
	"alfred/modules/UserProfileService/models"
	"alfred/modules/UserProfileService/pb"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type UserProfileService struct {
	fx.In
	db       *gorm.DB
	helloCli hellopb.HelloWorldServiceClient
}

//todo : AlWays add migration code for best practices
func NewUserProfileService(db *gorm.DB, helloCli hellopb.HelloWorldServiceClient) pb.UserProfileServiceServer {

	//initial migration of databases: schema migration
	models.InitialMigrationUserProfile(db)
	return &UserProfileService{
		db:       db,
		helloCli: helloCli,
	}
}
