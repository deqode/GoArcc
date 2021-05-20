package controllers

import (
	model "alfred/modules/account/v1/models"
	"context"
	"google.golang.org/genproto/protobuf/field_mask"
)

type AccountStore interface {
	//CreateAccount create new account details
	CreateAccount(ctx context.Context, in []*model.Account) ([]string, error)
	//GetAccount get account details by its unique id
	GetAccount(ctx context.Context, condition map[string]string) (*model.Account, error)
	//ListAccount get account details by its unique id
	ListAccount(ctx context.Context, condition map[string]string) ([]*model.Account, error)
	//UpdateAccount update existing account details
	UpdateAccount(ctx context.Context, in *model.Account, updateMask *field_mask.FieldMask) (*model.Account, error)
	//DeleteAccount delete the existing account from DB
	//This will be soft delete in DB
	DeleteAccount(ctx context.Context, condition map[string]string) (bool, error)
}
