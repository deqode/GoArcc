package AccountService

import (
	"alfred/modules/AccountService/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *AccountService) CreateAccount(context.Context, *pb.CreateAccountRequest) (*pb.Account, error){
	return nil, nil
}
func (s *AccountService) UpdateAccount(context.Context, *pb.UpdateAccountRequest) (*pb.Account, error){
	return nil, nil
}
func (s *AccountService) GetAccount(context.Context, *pb.GetAccountRequest) (*pb.Account, error){
	return nil, nil
}
func (s *AccountService) DeleteAccount(context.Context, *pb.DeleteAccountRequest) (*empty.Empty, error){
	return nil, nil
}
