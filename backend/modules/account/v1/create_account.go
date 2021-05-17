package account

import (
	"alfred/modules/account/v1/pb"
	"context"
)

func (s *accountsServer) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
	return s.store.CreateAccount(ctx, in)
}
