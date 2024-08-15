package health

import (
	"github/elliot9/ginExample/internal/pkg/context"
)

type Handler interface {
	Ping() context.HandlerFunc
}

type handler struct{}

func New() Handler {
	return &handler{}
}
