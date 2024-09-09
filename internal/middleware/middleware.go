package middleware

import (
	"github.com/elliot9/gin-example/internal/pkg/context"

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
