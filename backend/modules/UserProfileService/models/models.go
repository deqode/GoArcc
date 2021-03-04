package models

import (
	"alfred/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserProfile struct {
	Id          string
	Name        string
	Email       string
	PhoneNumber string
	UserId      string
	Source      string
}

func InitialMigrationUserProfile(db *gorm.DB) {
	if err := db.AutoMigrate(&UserProfile{}); err != nil {
		logger.Log.Debug("unable to migrate user profile service", zap.Error(err))
	}
	//Check if table exist or not
	if !db.Migrator().HasTable(&UserProfile{}) {
		// if table not exist then create table
		if err := db.Migrator().CreateTable(&UserProfile{}); err != nil {
			logger.Log.Debug("unable to create table")
		}
	}
}
