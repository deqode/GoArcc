package models

import (
	"alfred.sh/common/logger"
	"alfred/modules/stack/v1/pb"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type Stack struct {
	ID string `gorm:"primarykey"`
	// name of the stack
	Name string
	// slug of the stack
	Slug string
	// account id of stack
	AccountId string
	// git url
	GitURL string
	// git branch
	GitBranch string
	// username of user vcs
	UserName string
	// dev or production
	StackEnvironment pb.StackEnvironment
	// stack type will be go, ruby, node etc
	StackType string
	// status of stack deployed, failed,
	Status string
	// url where stack is deployed
	WebURL string
	// archived date
	ArchivedAt time.Time
	// stack released date
	ReleasedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func InitialMigrationStack(db *gorm.DB) {
	if err := db.AutoMigrate(&Stack{}); err != nil {
		logger.Log.Debug("unable to migrate stack service", zap.Error(err))
	}
	//Check if table exist or not
	if !db.Migrator().HasTable(&Stack{}) {
		// if table not exist then create table
		if err := db.Migrator().CreateTable(&Stack{}); err != nil {
			logger.Log.Debug("unable to create Stack table")
		}
	}
}

func (o *Stack) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}
