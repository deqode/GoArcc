package internal_svc

import (
	model "alfred/modules/account/v1/models"
	"alfred/modules/account/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s accountsIntServer) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}

	//prepare insert object
	accountModel := &model.Account{
		Slug:   in.Account.Slug,
		UserID: in.Account.UserId,
	}
	//insert into db
	if s.db.Create(accountModel).Error != nil {
		return nil, status.Error(codes.Internal, "Record not inserted in Account")
	}
	return &pb.Account{
		Id:     accountModel.ID,
		UserId: accountModel.UserID,
		Slug:   accountModel.Slug,
	}, nil
}
