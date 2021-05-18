package account

import (
	"alfred/modules/account/v1/pb"
	"context"
)

func (s *accountsInServer) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	return s.store.CreateAccount(ctx, in)
}
