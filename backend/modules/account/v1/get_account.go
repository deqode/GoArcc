package account

import (
	"alfred/modules/account/v1/pb"
	"context"
)

func (s *accountsServer) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error) {
	return s.store.GetAccount(ctx, in)
}
