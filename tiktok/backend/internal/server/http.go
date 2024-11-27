package server

import (
	v1 "backend/api/helloworld/v1"
	userV1 "backend/api/v1"
	"backend/internal/conf"
	"backend/internal/server/middleware/auth"
	"backend/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, s *conf.Security, greeter *service.GreeterService, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Address(c.Http.Network),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:3000"}),
			handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Origin"}),
			handlers.AllowCredentials(),
		)),
		http.Middleware(
			auth.NewJWTMiddleware(s.Jwt.Secret),
			recovery.Recovery(),
			tracing.Server(), // 添加链路追踪中间件
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	userV1.RegisterUserHTTPServer(srv, user)
	return srv
}
