package middleware

import (
	"time"

	"github.com/elliot9/gin-example/internal/pkg/context"

	"go.uber.org/zap"
)

type LoggerMiddleware struct {
	logger *zap.Logger
}

func NewLoggerMiddleware(logger *zap.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{logger: logger}
}

func (m *LoggerMiddleware) Handle(next context.HandlerFunc) context.HandlerFunc {
	return func(c context.Context) {
		start := time.Now()
		next(c)
		m.logger.Info("Request processed",
			zap.String("method", c.Method()),
			zap.String("path", c.URI()),
			zap.Duration("duration", time.Since(start)),
		)
	}
}
