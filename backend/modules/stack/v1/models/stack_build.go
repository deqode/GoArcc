package models

import (
	"alfred.sh/common/logger"
	"alfred/modules/stack/v1/pb"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type StackBuild struct {
	ID      string `gorm:"primarykey"`
	StackID string
	// slug of the stack
	Slug string
	// account id of stack
	AccountID string
	// status of stack deployed, failed,
	StackBuildStatus pb.StackBuildStatus
	// url where stack is deployed
	LogStreamChannel string
	// Complete date
	CompletedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func InitialMigrationStackBuild(db *gorm.DB) {
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

func (o *StackBuild) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}

// PrepareStackBuildModel convert request object into model object
func PrepareStackBuildModel(sb *pb.StackBuild) *StackBuild {
	return &StackBuild{
		StackID:          sb.GetStackId(),
		Slug:             sb.GetSlug(),
		AccountID:        sb.GetAccountId(),
		StackBuildStatus: sb.GetStatus(),
		LogStreamChannel: sb.GetLogStreamChannel(),
	}
}

// PrepareStackBuildObj convert request object into model object
func PrepareStackBuildObj(sb *StackBuild) *pb.StackBuild {
	return &pb.StackBuild{
		Id:               sb.ID,
		StackId:          sb.StackID,
		Status:           sb.StackBuildStatus,
		Slug:             sb.Slug,
		AccountId:        sb.AccountID,
		LogStreamChannel: sb.LogStreamChannel,
		CompletedAt:      nil,
	}
}
