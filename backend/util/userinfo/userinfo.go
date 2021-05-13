package userinfo

import (
	"context"
	"time"
)

type UserInfo struct {
	ID          string
	Email       string
	Sub         string
	TokenExpiry time.Time
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

// FromContext returns the User value stored in ctx, if any.
func FromContext(ctx context.Context) UserInfo {
	if u, ok := ctx.Value(userKey).(UserInfo); ok {
		return u
	}

	return UserInfo{}
}

func NewContext(ctx context.Context, u UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
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
