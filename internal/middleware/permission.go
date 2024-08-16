package middleware

import (
	"github/elliot9/ginExample/internal/pkg/context"
)

type PermissionMiddleware struct{}

func NewPermissionMiddleware() *PermissionMiddleware {
	return &PermissionMiddleware{}
}

func (m *PermissionMiddleware) Handle(next context.HandlerFunc) context.HandlerFunc {
	return func(c context.Context) {
		next(c)
	}
}
