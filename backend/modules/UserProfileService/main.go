package UserProfileService

import (
	hellopb "alfred/modules/HelloWorldService/pb"
	"alfred/modules/UserProfileService/models"
	"alfred/modules/UserProfileService/pb"
	"gorm.io/gorm"
)

type UserProfileService struct {
	db       *gorm.DB
	helloCli hellopb.HelloWorldServiceClient
}

func NewUserProfileService(db *gorm.DB, helloCli hellopb.HelloWorldServiceClient) pb.UserProfileServiceServer {
	//initial migration of databases: schema migration
	models.InitialMigrationUserProfile(db)
	return &UserProfileService{
		db:       db,
		helloCli: helloCli,
	}
}
