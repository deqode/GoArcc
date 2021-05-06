package Account

import (
	"alfred/modules/Account/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *accountsServer) GetUserAccounts(ctx context.Context, in *pb.GetUserAccountsRequest) (*pb.GetUserAccountsResponse, error) {
	var accounts []*pb.Account
	result := s.db.First(&accounts, "user_id = ?", in.UserId)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}
	return &pb.GetUserAccountsResponse{
		Accounts: accounts,
	}, nil
}
