package middleware

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
)

func Logging(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			startTime := time.Now()

			reply, err = handler(ctx, req)

			// 记录请求日志
			level := log.LevelInfo
			if err != nil {
				level = log.LevelError
			}

			_ = logger.Log(level,
				"method", ctx.Value("operation"),
				"latency", time.Since(startTime).Seconds(),
				"error", err,
			)

			return
		}
	}
}
