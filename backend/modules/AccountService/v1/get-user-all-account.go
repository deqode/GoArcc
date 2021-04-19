package AccountService

import (
	"alfred/modules/AccountService/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AccountService) GetUserAllAccount(ctx context.Context, in *pb.GetUserAllAccountRequest) (*pb.GetUserAllAccountResponse, error) {
	var accounts []*pb.Account
	result := s.db.First(&accounts, "user_id = ?", in.UserId)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}
	return &pb.GetUserAllAccountResponse{
		Accounts: accounts,
	}, nil
}
