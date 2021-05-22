package external_svc

import (
	_ "alfred.sh/common/database/helper"
	model "alfred/modules/account/v1/models"
	"alfred/modules/account/v1/pb"
	"context"
	"errors"
	"gorm.io/gorm"
)

// GetAccount Information with id
func (s accountExtServer) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error) {
	//validate user
	if err := s.ValidateUser(ctx); err != nil {
		return nil, err
	}
	//request validation
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	//account model
	account := model.Account{}
	gormDb := s.db
	if err := gormDb.First(&account, "id = ?", in.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &pb.Account{
		Id:     account.ID,
		Slug:   account.Slug,
		UserId: account.UserID,
	}, nil
}
