package account

import (
	"alfred/modules/account/v1/pb"
	"context"
)

func (s *accountsServer) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	return s.store.GetAccount(ctx, in)
}
