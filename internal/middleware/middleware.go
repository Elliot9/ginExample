package middleware

import (
	"github/elliot9/ginExample/internal/pkg/context"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	Handle(next context.HandlerFunc) context.HandlerFunc
}

func AdaptMiddleware(m Middleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.NewContext(c)
		m.Handle(func(ctx context.Context) {
			c.Next()
		})(ctx)
	}
}
