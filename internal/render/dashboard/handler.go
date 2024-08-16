package dashboard

import (
	"github/elliot9/ginExample/internal/pkg/context"
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
