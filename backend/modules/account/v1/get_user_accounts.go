package account

import (
	"alfred/modules/account/v1/pb"
	"context"
)

func (s *accountsServer) GetUserAccounts(ctx context.Context, in *pb.GetUserAccountsRequest) (*pb.GetUserAccountsResponse, error) {
	return s.store.GetUserAccounts(ctx, in)
}
