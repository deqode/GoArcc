package models

import (
	"github.com/deqode/GoArcc/logger"
	"github.com/deqode/GoArcc/protos/types"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type UserProfile struct {
	ID            string `gorm:"primarykey"`
	Name          string
	UserName      string
	Email         string
	PhoneNumber   string
	Sub           string `gorm:"uniqueIndex"`
	ProfilePicURL string
	Source        types.VCSProviders
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
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
