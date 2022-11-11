package telemetry

import (
	"context"
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger() *zap.Logger {
	return instrumentation.Logger
}

// LoggerFromContext returns a logger from context if it exists,
// and returns the standard instrumentation logger otherwise.
func LoggerFromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(LoggerContext{}).(*zap.Logger)
	if !ok {
		logger = instrumentation.Logger
	}

	return logger
}

func SetLoggerToContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, LoggerContext{}, logger)
}

func NewLogger(w io.Writer) *zap.Logger {
	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(w),
			zapcore.InfoLevel,
		),
	)
}
