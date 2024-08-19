package article

import (
	"github/elliot9/ginExample/internal/pkg/context"
)

type Handler interface {
	CreatePage() context.HandlerFunc
}

type handler struct {
	// cache  redis.Repo
	// logger *zap.Logger
}

func New() Handler {
	return &handler{}
}
