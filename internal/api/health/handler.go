package health

import (
	"github.com/elliot9/gin-example/internal/pkg/context"
)

type Handler interface {
	Ping() context.HandlerFunc
}

type handler struct{}

func New() Handler {
	return &handler{}
}
