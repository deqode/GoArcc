package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UserService) UserService

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserService) UserService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateUser(ctx context.Context, request CreateUserRequest) (resp *CreateUserResponse, err error) {
	defer func() {
		l.logger.Log("method", "CreateUser", "request", request, "resp", resp, "err", err)
	}()
	return l.next.CreateUser(ctx, request)
}
func (l loggingMiddleware) UpdateUser(ctx context.Context, request UpdateUserRequest) (resp *UpdateUserResponse, err error) {
	defer func() {
		l.logger.Log("method", "UpdateUser", "request", request, "resp", resp, "err", err)
	}()
	return l.next.UpdateUser(ctx, request)
}
func (l loggingMiddleware) GetUser(ctx context.Context, request GetUserRequest) (resp *GetUserResponse, err error) {
	defer func() {
		l.logger.Log("method", "GetUser", "request", request, "resp", resp, "err", err)
	}()
	return l.next.GetUser(ctx, request)
}
func (l loggingMiddleware) DeleteUser(ctx context.Context, request DeleteUserRequest) (resp *DeleteUserResponse, err error) {
	defer func() {
		l.logger.Log("method", "DeleteUser", "request", request, "resp", resp, "err", err)
	}()
	return l.next.DeleteUser(ctx, request)
}
