package models

import (
	"alfred/logger"
	"alfred/protos/types"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type VCSConnection struct {
	ID                 string
	Provider           types.GitProviders
	ConnectionID       string
	AccessToken        string
	RefreshToken       string
	AccessTokenExpiry  *time.Time
	RefreshTokenExpiry *time.Time
	Revoked            bool
	AccountID          string
	UserName           string
	Label              string
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
			logger.Log.Debug("unable to create vcs-connection table")
		}
	}
}
