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
	ID string `gorm:"type:uuid;primarykey"`
	// name of the stack
	Name string
	// slug of the stack
	Slug string
	// account id of stack
	AccountID string `gorm:"type:uuid"`
	// cloud_connection_id required to identify cloud provider
	CloudConnectionID string `gorm:"type:uuid"`
	// vcs_connection_id of user vcs
	VCSConnectionID string `gorm:"type:uuid"`
	// git url
	GitURL string
	// git branch
	GitBranch string
	// dev or production
	StackEnvironment pb.StackEnvironment
	// stack type will be go, ruby, node etc
	StackType pb.StackType
	// status of stack deployed, failed,
	Status pb.StackStatus
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

// PrepareStackModel convert request object into model object
func PrepareStackModel(stack *pb.Stack) *Stack {
	return &Stack{
		Name:              stack.GetName(),
		Slug:              stack.GetSlug(),
		AccountID:         stack.GetAccountId(),
		CloudConnectionID: stack.GetCloudConnectionId(),
		VCSConnectionID:   stack.GetVcsConnectionId(),
		GitURL:            stack.GetGitUrl(),
		GitBranch:         stack.GetGitBranch(),
		StackEnvironment:  stack.GetEnvironment(),
		StackType:         stack.GetStackType(),
		Status:            stack.GetStatus(),
		WebURL:            stack.GetWebUrl(),
	}
}

// PrepareStackObj convert model object into request object
func PrepareStackObj(stack *Stack) *pb.Stack {

	return &pb.Stack{
		Id:                stack.ID,
		Name:              stack.Name,
		Slug:              stack.Slug,
		AccountId:         stack.AccountID,
		CloudConnectionId: stack.CloudConnectionID,
		VcsConnectionId:   stack.VCSConnectionID,
		GitUrl:            stack.GitURL,
		GitBranch:         stack.GitBranch,
		Environment:       stack.StackEnvironment,
		StackType:         stack.StackType,
		//ArchivedAt:        stack.ArchivedAt,
		//ReleasedAt:        stack.ReleasedAt,
		Status: stack.Status,
	}
}

//TODO - time related conversion
