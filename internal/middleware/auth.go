package middleware

import (
	"github/elliot9/ginExample/internal/pkg/context"
	"net/http"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next context.HandlerFunc) context.HandlerFunc {
	return func(c context.Context) {
		if user := c.Session().Get(context.SessionAuthKey); user == nil {
			c.Redirect(http.StatusFound, "/admin/login")
			return
		}

		next(c)
	}
}
