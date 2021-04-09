package models

import (
	"alfred/logger"
	"alfred/modules/UserProfileService/v1/pb"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type UserProfile struct {
	ID          string
	Name        string
	Email       string
	PhoneNumber string
	Sub         string
	Source      pb.SOURCE
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func InitialMigrationUserProfile(db *gorm.DB) {
	if err := db.AutoMigrate(&UserProfile{}); err != nil {
		logger.Log.Debug("unable to migrate user profile service", zap.Error(err))
	}
	//Check if table exist or not
	if !db.Migrator().HasTable(&UserProfile{}) {
		// if table not exist then create table
		if err := db.Migrator().CreateTable(&UserProfile{}); err != nil {
			logger.Log.Debug("unable to create UserProfile table")
		}
	}
}
