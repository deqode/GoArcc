package UserProfileService

import (
	"alfred/modules/UserProfileService/models"
	"alfred/modules/UserProfileService/pb"
	"gorm.io/gorm"
)

type UserProfileService struct {
	db *gorm.DB
}

func NewUserProfileService(db *gorm.DB) pb.UserProfileServiceServer {
	//initial migration of databases: schema migration
	models.InitialMigrationUserProfile(db)
	return &UserProfileService{
		db: db,
	}
}
