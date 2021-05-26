package models

import (
	"alfred.sh/common/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

// Account - A user can have multiple accounts
type Account struct {
	ID        string `gorm:"primary_key;unique;type:uuid"`
	Slug      string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// InitialMigrationAccount will create table in the database
func InitialMigrationAccount(db *gorm.DB) {
	if err := db.AutoMigrate(&Account{}); err != nil {
		logger.Log.Debug("unable to migrate account service", zap.Error(err))
	}
	//Check if table exist or not
	if !db.Migrator().HasTable(&Account{}) {
		// if table not exist then create table
		if err := db.Migrator().CreateTable(&Account{}); err != nil {
			logger.Log.Debug("unable to account table")
		}
	}
}

func (o *Account) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}
