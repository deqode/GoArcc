package models

import (
	"alfred/modules/account/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/go-uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type accountPgStore struct {
	db *gorm.DB
}

func NewAccountStore(db *gorm.DB) AccountStore {
	return &accountPgStore{db}
}

func (s *accountPgStore) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}
	account := &Account{
		ID:     id,
		Slug:   in.GetAccount().GetSlug(),
		UserID: in.GetAccount().GetUserId(),
	}
	t := s.db.Create(account)
	if t.Error != nil {
		return nil, t.Error
	}
	in.Account.Id = id
	return in.Account, nil
}

func (s *accountPgStore) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error) {
	account := pb.Account{}
	err := s.db.First(&account, "id = ?", in.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err.Error != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &account, nil
}

func (s *accountPgStore) GetUserAccounts(ctx context.Context, in *pb.GetUserAccountsRequest) (*pb.GetUserAccountsResponse, error) {
	var accounts []*pb.Account
	result := s.db.First(&accounts, "user_id = ?", in.UserId)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}
	return &pb.GetUserAccountsResponse{
		Accounts: accounts,
	}, nil
}

func (s *accountPgStore) UpdateAccount(ctx context.Context, in *pb.UpdateAccountRequest) (*pb.Account, error) {
	account := &Account{
		Slug: in.GetAccount().GetSlug(),
	}
	t := s.db.Model(&account).Where("id ="+in.Account.Id, true).Update("slug", in.Account.Slug)
	if t.Error != nil {
		return nil, t.Error
	}
	return in.Account, nil
}

func (s *accountPgStore) DeleteAccount(context.Context, *pb.DeleteAccountRequest) (*empty.Empty, error) {
	return nil, nil
}
