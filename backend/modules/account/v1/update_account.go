package account

import (
	"alfred/modules/account/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *accountsServer) UpdateAccount(ctx context.Context, in *pb.UpdateAccountRequest) (*pb.Account, error) {
	if in.Account == nil {
		return nil, status.Error(codes.FailedPrecondition, "Account to update is not provided")
	}
	if len(in.GetAccount().GetId()) < 3 {
		return nil, status.Error(codes.FailedPrecondition, "Account Id is not provided")
	}
	return s.store.UpdateAccount(ctx, in)
}
