package middleware

import (
	"github/elliot9/ginExample/internal/auth"
	"github/elliot9/ginExample/internal/pkg/context"
	"net/http"
)

type GuestMiddleware struct{}

func NewGuestMiddleware() *GuestMiddleware {
	return &GuestMiddleware{}
}

func (m *GuestMiddleware) Handle(next context.HandlerFunc) context.HandlerFunc {
	return func(c context.Context) {
		if user := auth.New().Me(c); user != nil {
			c.Redirect(http.StatusFound, "/admin")
			return
		}

		next(c)
	}
}
