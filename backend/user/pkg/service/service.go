package service

import "context"

// UserService describes the service.
type UserService interface {
	CreateUser(ctx context.Context, request CreateUserRequest) (resp *CreateUserResponse, err error)
	UpdateUser(ctx context.Context, request UpdateUserRequest) (resp *UpdateUserResponse, err error)
	GetUser(ctx context.Context, request GetUserRequest) (resp *GetUserResponse, err error)
	DeleteUser(ctx context.Context, request DeleteUserRequest) (resp *DeleteUserResponse, err error)
}

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Source      string `json:"source"`
}

type ErrorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CreateUserRequest struct {
	User *User `json:"user"`
}

type CreateUserResponse struct {
	User *User         `json:"user"`
	Err  *ErrorMessage `json:"err"`
}

type UpdateUserRequest struct {
	User *User `json:"user"`
}

type UpdateUserResponse struct {
	User *User         `json:"user"`
	Err  *ErrorMessage `json:"err"`
}

type GetUserRequest struct {
	Id string `json:"id"`
}
type GetUserResponse struct {
	User *User `json:"user"`
}

type DeleteUserRequest struct {
	Id string `json:"id"`
}

type DeleteUserResponse struct {
	Err *ErrorMessage `json:"err"`
}

type basicUserService struct{}

func (b *basicUserService) CreateUser(ctx context.Context, request CreateUserRequest) (resp *CreateUserResponse, err error) {
	res := &CreateUserResponse{
		User: &User{
		 Name: "Hello sir !",
		},
	}
	return res, err
}
func (b *basicUserService) UpdateUser(ctx context.Context, request UpdateUserRequest) (resp *UpdateUserResponse, err error) {
	// TODO implement the business logic of UpdateUser
	return resp, err
}
func (b *basicUserService) GetUser(ctx context.Context, request GetUserRequest) (resp *GetUserResponse, err error) {
	// TODO implement the business logic of GetUser
	return resp, err
}
func (b *basicUserService) DeleteUser(ctx context.Context, request DeleteUserRequest) (resp *DeleteUserResponse, err error) {
	// TODO implement the business logic of DeleteUser
	return resp, err
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService() UserService {
	return &basicUserService{}
}

// New returns a UserService with all of the expected middleware wired in.
func New(middleware []Middleware) UserService {
	var svc UserService = NewBasicUserService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
