package external_svc

import (
	databaseHelper "alfred.sh/common/database/helper"
	model "alfred/modules/account/v1/models"
	"alfred/modules/account/v1/pb"
	"alfred/util/userinfo"
	"context"
)

// GetAccount Information with id
func (s accountExtServer) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error) {

	userId := userinfo.FromContext(ctx).Sub
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	//account model
	account := model.Account{}
	gormDb := s.db
	//ie: Select * from account where id = in.id
	tx := gormDb.Where(&account, in.Id).Where("user_id = ", userId)
	if err := databaseHelper.ValidateResult(tx); err != nil {
		return nil, err
	}
	return &pb.Account{
		Id:     account.ID,
		Slug:   account.Slug,
		UserId: account.UserID,
	}, nil
}
