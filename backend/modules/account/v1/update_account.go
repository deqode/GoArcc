package account

import (
	"alfred/modules/account/v1/pb"
	"context"
)

func (s *accountsServer) UpdateAccount(ctx context.Context, in *pb.UpdateAccountRequest) (*pb.Account, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	return s.store.UpdateAccount(ctx, in)
}
