package auth

import (
	"backend/internal/pkg/auth"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v4"
)

func NewJWTMiddleware(key string) middleware.Middleware {
	return selector.Server(
		beforeJWT(),
		jwt.Server(
			func(token *jwt2.Token) (interface{}, error) {
				return []byte(key), nil
			},
			jwt.WithSigningMethod(jwt2.SigningMethodHS256),
			jwt.WithClaims(func() jwt2.Claims { return &auth.Claims{} }),
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

func beforeJWT() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := tr.(http.Transporter); ok {
					// 从 cookie 中获取 token
					cookie, err := ht.Request().Cookie("user_token")
					if err != nil {
						return nil, errors.Unauthorized("AUTH_ERROR", "未登录")
					}

					// 将 token 添加到 Authorization header
					ht.RequestHeader().Set("Authorization", fmt.Sprintf("Bearer %s", cookie.Value))
				}
			}
			return handler(ctx, req)
		}
	}
}
