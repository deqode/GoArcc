package AccountService

import (
	"alfred/modules/AccountService/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *AccountService) DeleteAccount(context.Context, *pb.DeleteAccountRequest) (*empty.Empty, error) {
	return nil, nil
}
