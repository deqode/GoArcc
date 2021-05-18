package models

import (
	"alfred/modules/account/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

type AccountStore interface {
	//CreateAccount create new account details
	CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error)
	//UpdateAccount update existing account details
	UpdateAccount(ctx context.Context, in *pb.UpdateAccountRequest) (*pb.Account, error)
	//GetAccount get account details by its unique id
	GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error)
	//GetUserAccounts will list all the accounts of user
	GetUserAccounts(ctx context.Context, in *pb.GetUserAccountsRequest) (*pb.GetUserAccountsResponse, error)
	//DeleteAccount delete the existing account from DB
	//This will be soft delete in DB
	DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*empty.Empty, error)
}
