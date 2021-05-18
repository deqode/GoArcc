package account

import (
	"alfred/modules/account/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *accountsServer) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*empty.Empty, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	return s.store.DeleteAccount(ctx, in)
}
