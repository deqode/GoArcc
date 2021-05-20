package internal_svc

import (
	"alfred/modules/account/v1/internal-svc/pb"
	model "alfred/modules/account/v1/models"
	"context"
	"gorm.io/gorm"
	"time"
)

func (s accountsIntServer) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	var accounts []*model.Account
	accountModel := &model.Account{
		Slug:      in.Account.Slug,
		UserID:    in.Account.UserId,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
	accounts = append(accounts, accountModel)
	res, err := s.store.CreateAccount(ctx, accounts)
	if err != nil {
		return nil, err
	}
	return &pb.Account{
		Id: res[0],
	}, nil
}
