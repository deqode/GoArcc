package models

import (
	"alfred/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CloudConnection struct {
	ID           string
	ClientId     string
	ClientSecret string
	AccountId    string
	Scope        []string
	Provider     string
}

func InitialMigrationCloudConnection(db *gorm.DB) {
	if err := db.AutoMigrate(&CloudConnection{}); err != nil {
		logger.Log.Debug("unable to migrate user profile service", zap.Error(err))
	}
	//Check if table exist or not
	if !db.Migrator().HasTable(&CloudConnection{}) {
		// if table not exist then create table
		if err := db.Migrator().CreateTable(&CloudConnection{}); err != nil {
			logger.Log.Debug("unable to create table")
		}
	}
}
