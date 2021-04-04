package models

import (
	"alfred/logger"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VCSConnection struct {
	ID                 string
	Provider           int32
	ConnectionId       string
	AccessToken        string
	RefreshToken       string
	AccessTokenExpiry  *timestamp.Timestamp
	RefreshTokenExpiry *timestamp.Timestamp
	Revoked            bool
	AccountId          string
}

func InitialMigrationVCSConnection(db *gorm.DB) {
	if err := db.AutoMigrate(&VCSConnection{}); err != nil {
		logger.Log.Debug("unable to migrate user profile service", zap.Error(err))
	}
	//Check if table exist or not
	if !db.Migrator().HasTable(&VCSConnection{}) {
		// if table not exist then create table
		if err := db.Migrator().CreateTable(&VCSConnection{}); err != nil {
			logger.Log.Debug("unable to create table")
		}
	}
}
