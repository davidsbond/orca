package log

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

type (
	ctxKey struct{}
)

func ToContext(ctx context.Context, log *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, log)
}

func FromContext(ctx context.Context) *slog.Logger {
	log, ok := ctx.Value(ctxKey{}).(*slog.Logger)
	if !ok {
		return slog.Default()
	}

	return log
}

func Interceptor(ctx context.Context) logging.Logger {
	logger := FromContext(ctx)

	return logging.LoggerFunc(func(ctx context.Context, _ logging.Level, msg string, fields ...any) {
		logger.Log(ctx, slog.LevelDebug, msg, fields...)
	})
}
