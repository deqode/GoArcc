package userinfo

import (
	databaseHelper "alfred.sh/common/database/helper"
	"alfred/modules/user-profile/v1/models"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	ID          string
	Email       string
	Sub         string
	TokenExpiry time.Time
}

type ValidateUserInfo struct {
	Ctx          context.Context
	RootTable    interface{}
	RootTableTag string
	//Key Will be tag name
	Args               map[string]interface{}
	SkipRootValidation bool
	SkipArgsValidation bool
	Db                 *gorm.DB
	skipValidation     bool
}

const (
	userKey key = iota
	clientIpKey
	grpcCallKey
)

type key int

func GetClientIP(ctx context.Context) string {
	if ip, ok := ctx.Value(clientIpKey).(string); ok {
		return ip
	}
	return ""
}

func NewContextWithClientIP(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, clientIpKey, ip)
}

func NewContextForGrpcCall(ctx context.Context) context.Context {
	return context.WithValue(ctx, grpcCallKey, true)
}

func IsGrpcCall(ctx context.Context) bool {
	if v, ok := ctx.Value(grpcCallKey).(bool); ok {
		return v
	}
	return false
}

func WithClaims(ctx context.Context, c map[string]interface{}) context.Context {
	return context.WithValue(ctx, userKey, c)
}

// FromContext returns the User value stored in ctx, if any.
func FromContext(ctx context.Context) (usr UserInfo) {
	usr = FromClaims(ctx.Value(userKey).(map[string]interface{}))
	return
}

func NewContext(ctx context.Context, u UserInfo) context.Context {
	m := make(map[string]interface{})
	m["sub"] = u.Sub
	return context.WithValue(ctx, userKey, m)
}

func FromClaims(claims map[string]interface{}) (ui UserInfo) {
	if claims == nil {
		return
	}

	if v, ok := claims["sub"]; ok {
		ui.Sub = v.(string)
		ui.ID = v.(string)
	}
	if v, ok := claims["ext"]; ok {
		if e, exist := v.(map[string]interface{})["email"]; exist {
			ui.Email = e.(string)
		}
	}

	if v, ok := claims["exp"]; ok {
		tm := time.Unix(int64(v.(float64)), 0)
		ui.TokenExpiry = tm
	}

	return
}

func (validateUserInfo *ValidateUserInfo) validate() error {
	if validateUserInfo == nil {
		return status.Error(codes.FailedPrecondition, "authentication validation failed")
	}
	if !validateUserInfo.SkipRootValidation && (validateUserInfo.RootTableTag == "" || validateUserInfo.RootTable == nil) {
		return status.Error(codes.FailedPrecondition, "authentication validation failed")
	}
	if !validateUserInfo.SkipArgsValidation && validateUserInfo.Args == nil {
		return status.Error(codes.FailedPrecondition, "authentication validation failed")
	}
	return nil
}

// BasicAuthValidation It will check the user id present in context is valid or not
func BasicAuthValidation(ctx context.Context, db *gorm.DB) error {
	v := &ValidateUserInfo{
		Ctx:                ctx,
		RootTable:          &models.UserProfile{},
		RootTableTag:       "id",
		Args:               nil,
		SkipRootValidation: false,
		SkipArgsValidation: true,
		Db:                 db,
		skipValidation:     true,
	}
	if err := v.ValidateUser(); err != nil {
		return err
	}
	return nil
}

// ValidateUser Validate User : Will validate user id is present in the given table or not
func (validateUserInfo *ValidateUserInfo) ValidateUser() error {
	if !validateUserInfo.skipValidation {
		if err := validateUserInfo.validate(); err != nil {
			return err
		}
	}
	//Extract user id from context
	userInfo := FromContext(validateUserInfo.Ctx)
	if userInfo.Sub == "" {
		return status.Error(codes.PermissionDenied, "unauthenticated user")
	}
	if !validateUserInfo.SkipRootValidation {
		tx := validateUserInfo.Db.Where(validateUserInfo.RootTableTag+" = ?", userInfo.ID).First(validateUserInfo.RootTable)
		if err := databaseHelper.ValidateResult(tx); err != nil {
			return status.Error(codes.PermissionDenied, "unauthenticated user")
		}
	}

	if !validateUserInfo.SkipArgsValidation {
		for tag, argTable := range validateUserInfo.Args {
			tx := validateUserInfo.Db.Where(tag+" = ?", userInfo.ID).First(argTable)
			if err := databaseHelper.ValidateResult(tx); err != nil {
				return status.Error(codes.PermissionDenied, "unauthenticated user")
			}
		}
	}
	return nil
}
