package account

import (
	"alfred/modules/account/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *accountsServer) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*empty.Empty, error) {
	return s.store.DeleteAccount(ctx, in)
}
