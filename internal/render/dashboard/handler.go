package dashboard

import (
	"github.com/elliot9/gin-example/internal/pkg/context"
)

type Handler interface {
	IndexPage() context.HandlerFunc
}

type handler struct {
	// cache  redis.Repo
	// logger *zap.Logger
}

func New() Handler {
	return &handler{}
}
