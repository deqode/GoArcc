package external_svc

import (
	pb "alfred/modules/account/v1/external-svc/pb"
	"context"
)

func (s accountExtServer) GetUserAccounts(ctx context.Context, in *pb.GetUserAccountsRequest) (*pb.GetUserAccountsResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	m["user_id"] = in.UserId
	acc, err := s.store.ListAccount(ctx, m)
	if err != nil {
		return nil, err
	}
	var accounts []*pb.Account
	for _, v := range acc {
		accounts = append(accounts, &pb.Account{
			Id:     v.ID,
			Slug:   v.Slug,
			UserId: v.UserID,
		})
	}
	return &pb.GetUserAccountsResponse{
		Accounts: accounts,
	}, nil
}
