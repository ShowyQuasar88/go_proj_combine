package main

import (
	"context"
	"flag"
	"os"

	"backend/internal/conf"

	sys_log "backend/internal/pkg/log"
	"backend/internal/pkg/trace"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "tiktok-backend"
	// Version is the version of the compiled software.
	Version string = "1.0.0"
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

// initLogger 初始化日志
func initLogger(c *conf.Log) log.Logger {
	return sys_log.NewLogger(&sys_log.Options{
		Level:      c.Level,
		Filename:   c.Filename,
		MaxSize:    int(c.MaxSize),
		MaxAge:     int(c.MaxAge),
		MaxBackups: int(c.MaxBackups),
	})
}

// initTrace 初始化链路追踪
func initTrace(c *conf.Trace) (*tracesdk.TracerProvider, error) {
	return trace.NewTracerProvider(&trace.Options{
		ServiceName: Name,
		Endpoint:    c.Endpoint,
		Sampler:     c.Sampler,
	})
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// 初始化日志
	logger := initLogger(bc.Log)

	// 初始化链路追踪
	tp, err := initTrace(bc.Trace)
	if err != nil {
		panic(err)
	}
	defer tp.Shutdown(context.Background())

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Security, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
