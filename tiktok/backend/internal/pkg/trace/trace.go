package trace

import (
	_ "context"
	_ "errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	_ "go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	_ "time"
)

type Options struct {
	ServiceName string  `json:"serviceName"`
	Endpoint    string  `json:"endpoint"`
	Sampler     float64 `json:"sampler"`
}

func InitTracer(opts *Options) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(opts.Endpoint)))
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		// 设置采样率
		tracesdk.WithSampler(tracesdk.TraceIDRatioBased(opts.Sampler)),
		// 设置批量处理
		tracesdk.WithBatcher(exp),
		// 设置资源信息
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(opts.ServiceName),
			attribute.String("environment", "production"))),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}
