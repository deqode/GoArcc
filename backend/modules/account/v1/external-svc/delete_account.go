package external_svc

import (
	pb "alfred/modules/account/v1/external-svc/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s accountExtServer) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*empty.Empty, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	m["id"] = in.Id
	_, err = s.store.DeleteAccount(ctx, m)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
