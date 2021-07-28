package models

import (
	"github.com/deqode/GoArcc/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            string `gorm:"primarykey"`
	FirstName     string
	LastName      string
	UserName      string
	Email         string
	PhoneNumber   string
	ProfilePicURL string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func InitialMigrationUser(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}); err != nil {
		logger.Log.Debug("unable to migrate user service", zap.Error(err))
	}
	//Check if table exist or not
	if !db.Migrator().HasTable(&User{}) {
		// if table not exist then create table
		if err := db.Migrator().CreateTable(&User{}); err != nil {
			logger.Log.Debug("unable to create User table")
		}
	}
}
