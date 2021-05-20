package external_svc

import (
	pb "alfred/modules/account/v1/external-svc/pb"
	"context"
)

func (s accountExtServer) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	m["id"] = in.Id
	acc, err := s.store.GetAccount(ctx, m)
	if err != nil {
		return nil, err
	}
	return &pb.Account{
		Id:     acc.ID,
		Slug:   acc.Slug,
		UserId: acc.UserID,
	}, nil
}
