package models

import (
	"alfred.sh/common/logger"
	"alfred/protos/types"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type VCSConnection struct {
	ID                 string `gorm:"primary_key;unique;type:uuid"`
	Provider           types.VCSProviders
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

func (o *VCSConnection) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}
