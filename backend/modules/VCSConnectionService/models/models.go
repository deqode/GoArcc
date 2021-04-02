package models

import (
	"alfred/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VCSConnection struct {
	ID           string
	ClientId     string
	ClientSecret string
	AccountId    string
	Scope        string
	Provider     string
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
