package external_svc

import (
	"alfred/modules/account/v1/common"
	pb "alfred/modules/account/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s accountExtServer) UpdateAccount(ctx context.Context, in *pb.UpdateAccountRequest) (*pb.Account, error) {
	if in.Account == nil {
		return nil, status.Error(codes.FailedPrecondition, "Account to update is not provided")
	}
	if len(in.GetAccount().GetId()) < 3 {
		return nil, status.Error(codes.FailedPrecondition, "Account Id is not provided")
	}
	//Authentication check
	if err := common.ValidateUser(ctx, in.Account.Id, s.db); err != nil {
		return nil, err
	}

	return &pb.Account{
		Id:     "",
		Slug:   "",
		UserId: "",
	}, nil
}
