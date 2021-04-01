package models

import (
	"alfred/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Account struct {
	ID   string
	Slug string
}

func InitialMigrationAccount(db *gorm.DB) {
	if err := db.AutoMigrate(&Account{}); err != nil {
		logger.Log.Debug("unable to migrate user profile service", zap.Error(err))
	}
	//Check if table exist or not
	if !db.Migrator().HasTable(&Account{}) {
		// if table not exist then create table
		if err := db.Migrator().CreateTable(&Account{}); err != nil {
			logger.Log.Debug("unable to create table")
		}
	}
}
