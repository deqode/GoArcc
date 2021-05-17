package account

import (
	"alfred/modules/account/v1/pb"
	"context"
)

func (s *accountsServer) UpdateAccount(ctx context.Context, in *pb.UpdateAccountRequest) (*pb.Account, error) {
	return s.store.UpdateAccount(ctx, in)
}
