package models

import (
	"alfred/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type VCSConnection struct {
	ID                 string
	Provider           int
	ConnectionId       string
	AccessToken        string
	RefreshToken       string
	AccessTokenExpiry  *time.Time
	RefreshTokenExpiry *time.Time
	Revoked            bool
	AccountId          string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

func InitialMigrationVCSConnection(db *gorm.DB) {
	if err := db.AutoMigrate(&VCSConnection{}); err != nil {
		logger.Log.Debug("unable to migrate user profile service", zap.Error(err))
	}
	//Check if table exist or not
	if !db.Migrator().HasTable(&VCSConnection{}) {
		// if table not exist then create table
		if err := db.Migrator().CreateTable(&VCSConnection{}); err != nil {
			logger.Log.Debug("unable to create VCSConnection table")
		}
	}
}
