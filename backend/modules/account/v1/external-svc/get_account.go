package external_svc

import (
	_ "alfred.sh/common/database/helper"
	model "alfred/modules/account/v1/models"
	"alfred/modules/account/v1/pb"
	"alfred/util/userinfo"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
	if err := gormDb.First(&account, "id = ? AND user_id = ?", in.Id, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.Unauthenticated, "unauthorized user")
		}
		return nil, err
	}
	return &pb.Account{
		Id:     account.ID,
		Slug:   account.Slug,
		UserId: account.UserID,
	}, nil
}
