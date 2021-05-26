package external_svc

import (
	model "alfred/modules/account/v1/models"
	"alfred/modules/account/v1/pb"
	"alfred/util/userinfo"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s accountExtServer) GetUserAccounts(ctx context.Context, empty *empty.Empty) (*pb.GetUserAccountsResponse, error) {
	//Authentication check
	usr := userinfo.FromContext(ctx)
	//ie: select * from user where user_id = in.userId
	var accountModel []*model.Account
	tx := s.db.Where("user_id = ?", usr.ID).Find(&accountModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	//prepare response
	var accounts []*pb.Account
	for _, v := range accountModel {
		if v.UserID != usr.ID {
			return nil, status.Error(codes.PermissionDenied, "unauthenticated user")
		}
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
