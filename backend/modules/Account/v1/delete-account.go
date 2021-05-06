package Account

import (
	"alfred/modules/Account/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *accountsServer) DeleteAccount(context.Context, *pb.DeleteAccountRequest) (*empty.Empty, error) {
	return nil, nil
}
