package account

import (
	"alfred/modules/account/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *accountsServer) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error) {
	account := pb.Account{}
	result := s.db.First(&account, "id = ?", in.Id)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}
	return &account, nil
}
