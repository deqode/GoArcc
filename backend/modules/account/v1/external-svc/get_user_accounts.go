package external_svc

import (
	databaseHelper "alfred.sh/common/database/helper"
	model "alfred/modules/account/v1/models"
	"alfred/modules/account/v1/pb"
	"context"
)

func (s accountExtServer) GetUserAccounts(ctx context.Context, in *pb.GetUserAccountsRequest) (*pb.GetUserAccountsResponse, error) {
	//validate user
	if err := s.ValidateUser(ctx); err != nil {
		return nil, err
	}
	//Request Validation
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	//ie: select * from user where user_id = in.userId
	var accountModel []*model.Account
	gormDb := s.db
	tx := gormDb.Where("user_id = ?", in.UserId).Find(&accountModel)
	if err := databaseHelper.ValidateResult(tx); err != nil {
		return nil, err
	}
	//prepare response
	var accounts []*pb.Account
	for _, v := range accountModel {
		accounts = append(accounts, &pb.Account{
			Id:     v.ID,
			Slug:   v.Slug,
			UserId: v.UserID,
		})
	}
	return &pb.GetUserAccountsResponse{
		Accounts: accounts,
	}, nil
}
