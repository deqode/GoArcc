package account

import (
	"alfred/modules/account/v1/pb"
	"context"
)

func (s *accountsServer) GetUserAccounts(ctx context.Context, in *pb.GetUserAccountsRequest) (*pb.GetUserAccountsResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	return s.store.GetUserAccounts(ctx, in)
}
