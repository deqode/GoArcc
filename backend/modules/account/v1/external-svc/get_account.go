package external_svc

import (
	databaseHelper "alfred.sh/common/database/helper"
	model "alfred/modules/account/v1/models"
	"alfred/modules/account/v1/pb"
	"context"
)

// GetAccount Information with id
func (s accountExtServer) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error) {

	err := in.Validate()
	if err != nil {
		return nil, err
	}
	//account model
	account := model.Account{}
	gormDb := s.db
	//ie: Select * from account where id = in.id
	tx := gormDb.First(&account, in.Id)
	if err := databaseHelper.ValidateResult(tx); err != nil {
		return nil , err
	}
	return &pb.Account{
		Id:     account.ID,
		Slug:   account.Slug,
		UserId: account.UserID,
	}, nil
}
