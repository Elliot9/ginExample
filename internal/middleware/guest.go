package middleware

import (
	"github/elliot9/ginExample/internal/pkg/context"
)

type GuestMiddleware struct{}

func NewGuestMiddleware() *GuestMiddleware {
	return &GuestMiddleware{}
}

func (m *GuestMiddleware) Handle(next context.HandlerFunc) context.HandlerFunc {
	return func(c context.Context) {
		if user := c.Session().Get(context.SessionAuthKey); user != nil {
			c.Abort(context.Error(403, 100, "Unauthorized"))
			return
		}

		next(c)
	}
}
