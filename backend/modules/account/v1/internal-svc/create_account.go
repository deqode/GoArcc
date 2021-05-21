package internal_svc

import (
	databaseHelper "alfred.sh/common/database/helper"
	model "alfred/modules/account/v1/models"
	"alfred/modules/account/v1/pb"
	"context"
	"gorm.io/gorm"
	"time"
)

func (s accountsIntServer) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}
	//prepare insert object
	accountModel := &model.Account{
		Slug:      in.Account.Slug,
		UserID:    in.Account.UserId,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
	//insert into db
	gormDb := s.db
	tx := gormDb.Create(accountModel)
	if err := databaseHelper.ValidateResult(tx); err != nil {
		return nil , err
	}
	return &pb.Account{
		Id: accountModel.ID,
		UserId: accountModel.UserID,
		Slug: accountModel.Slug,
	}, nil
}

