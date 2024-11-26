package auth

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwt2 "github.com/golang-jwt/jwt/v4"
)

func NewJWTMiddleware(key string) middleware.Middleware {
	return selector.Server(
		jwt.Server(
			func(token *jwt2.Token) (interface{}, error) {
				return []byte(key), nil
			},
		),
	).Match(NewWhiteListMatcher()).Build()
}

// NewWhiteListMatcher 创建白名单匹配器
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := map[string]struct{}{
		"/v1.User/Register": {},
		"/v1.User/Login":    {},
	}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
