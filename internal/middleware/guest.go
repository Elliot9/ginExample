package middleware

import (
	"net/http"

	"github.com/elliot9/gin-example/internal/auth"
	"github.com/elliot9/gin-example/internal/pkg/context"
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
