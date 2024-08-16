package middleware

import (
	"github/elliot9/ginExample/internal/pkg/context"
	"net/http"
)

type ErrorMiddleware struct{}

func NewErrorMiddleware() *ErrorMiddleware {
	return &ErrorMiddleware{}
}

func (m *ErrorMiddleware) Handle(next context.HandlerFunc) context.HandlerFunc {
	return func(c context.Context) {
		next(c)

		if err := c.GetAbort(); err != nil {
			switch err.HTTPCode() {
			case http.StatusBadRequest:
				c.HTML("errors/400.html", err.Message())
			case http.StatusForbidden:
				c.HTML("errors/403.html", err.Message())
			case http.StatusUnauthorized:
				c.HTML("errors/401.html", err.Message())
			default:

			}
			return
		}
	}
}
