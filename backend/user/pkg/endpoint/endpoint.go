package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "user/pkg/service"
)

// CreateUserRequest collects the request parameters for the CreateUser method.
type CreateUserRequest struct {
	Request service.CreateUserRequest `json:"request"`
}

// CreateUserResponse collects the response parameters for the CreateUser method.
type CreateUserResponse struct {
	Resp *service.CreateUserResponse `json:"resp"`
	Err  error               `json:"err"`
}

// MakeCreateUserEndpoint returns an endpoint that invokes CreateUser on the service.
func MakeCreateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		resp, err := s.CreateUser(ctx, req.Request)
		return CreateUserResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateUserResponse) Failed() error {
	return r.Err
}

// UpdateUserRequest collects the request parameters for the UpdateUser method.
type UpdateUserRequest struct {
	Request service.UpdateUserRequest `json:"request"`
}

// UpdateUserResponse collects the response parameters for the UpdateUser method.
type UpdateUserResponse struct {
	Resp *service.UpdateUserResponse `json:"resp"`
	Err  error               `json:"err"`
}

// MakeUpdateUserEndpoint returns an endpoint that invokes UpdateUser on the service.
func MakeUpdateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		resp, err := s.UpdateUser(ctx, req.Request)
		return UpdateUserResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateUserResponse) Failed() error {
	return r.Err
}

// GetUserRequest collects the request parameters for the GetUser method.
type GetUserRequest struct {
	Request service.GetUserRequest `json:"request"`
}

// GetUserResponse collects the response parameters for the GetUser method.
type GetUserResponse struct {
	Resp *service.GetUserResponse `json:"resp"`
	Err  error            `json:"err"`
}

// MakeGetUserEndpoint returns an endpoint that invokes GetUser on the service.
func MakeGetUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		resp, err := s.GetUser(ctx, req.Request)
		return GetUserResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserResponse) Failed() error {
	return r.Err
}

// DeleteUserRequest collects the request parameters for the DeleteUser method.
type DeleteUserRequest struct {
	Request service.DeleteUserRequest `json:"request"`
}

// DeleteUserResponse collects the response parameters for the DeleteUser method.
type DeleteUserResponse struct {
	Resp *service.DeleteUserResponse `json:"resp"`
	Err  error               `json:"err"`
}

// MakeDeleteUserEndpoint returns an endpoint that invokes DeleteUser on the service.
func MakeDeleteUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		resp, err := s.DeleteUser(ctx, req.Request)
		return DeleteUserResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteUserResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateUser implements Service. Primarily useful in a client.
func (e Endpoints) CreateUser(ctx context.Context, request service.CreateUserRequest) (resp *service.CreateUserResponse, err error) {
	request0 := CreateUserRequest{Request: request}
	response, err := e.CreateUserEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response.(CreateUserResponse).Resp, response.(CreateUserResponse).Err
}

// UpdateUser implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUser(ctx context.Context, request service.UpdateUserRequest) (resp *service.UpdateUserResponse, err error) {
	request0 := UpdateUserRequest{Request: request}
	response, err := e.UpdateUserEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response.(UpdateUserResponse).Resp, response.(UpdateUserResponse).Err
}

// GetUser implements Service. Primarily useful in a client.
func (e Endpoints) GetUser(ctx context.Context, request service.GetUserRequest) (resp *service.GetUserResponse, err error) {
	request0 := GetUserRequest{Request: request}
	response, err := e.GetUserEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response.(GetUserResponse).Resp, response.(GetUserResponse).Err
}

// DeleteUser implements Service. Primarily useful in a client.
func (e Endpoints) DeleteUser(ctx context.Context, request service.DeleteUserRequest) (resp *service.DeleteUserResponse, err error) {
	request0 := DeleteUserRequest{Request: request}
	response, err := e.DeleteUserEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response.(DeleteUserResponse).Resp, response.(DeleteUserResponse).Err
}
