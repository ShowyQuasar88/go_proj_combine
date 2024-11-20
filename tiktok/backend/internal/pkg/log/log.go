package log

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Options struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxSize"`
	MaxAge     int    `json:"maxAge"`
	MaxBackups int    `json:"maxBackups"`
}

type zapLogger struct {
	log *zap.Logger
}

func (l *zapLogger) Log(level log.Level, keyVals ...interface{}) error {
	switch level {
	case log.LevelDebug:
		l.log.Debug(fmt.Sprint(keyVals...))
	case log.LevelInfo:
		l.log.Info(fmt.Sprint(keyVals...))
	case log.LevelWarn:
		l.log.Warn(fmt.Sprint(keyVals...))
	case log.LevelError:
		l.log.Error(fmt.Sprint(keyVals...))
	}
	return nil
}

// NewLogger 创建日志记录器，将 zap.Logger 转换为 kratos 的 Logger
func NewLogger(opts *Options) log.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 配置输出
	var cores []zapcore.Core

	// 控制台输出
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleCore := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(getLogLevel(opts.Level)),
	)
	cores = append(cores, consoleCore)

	// 文件输出
	if opts.Filename != "" {
		fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
		writer := getLogWriter(opts)
		fileCore := zapcore.NewCore(
			fileEncoder,
			zapcore.AddSync(writer),
			zap.NewAtomicLevelAt(getLogLevel(opts.Level)),
		)
		cores = append(cores, fileCore)
	}

	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &zapLogger{log: logger}
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func getLogWriter(opts *Options) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   opts.Filename,
		MaxSize:    opts.MaxSize,    // 每个文件最大尺寸，单位MB
		MaxBackups: opts.MaxBackups, // 保留的旧文件最大数量
		MaxAge:     opts.MaxAge,     // 保留的旧文件最大天数
		Compress:   true,            // 是否压缩
	}
}
