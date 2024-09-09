package dashboard

import (
	"github.com/elliot9/gin-example/internal/auth"
	"github.com/elliot9/gin-example/internal/pkg/context"

	"github.com/gin-gonic/gin"
)

func (h *handler) IndexPage() context.HandlerFunc {
	return func(c context.Context) {
		admin := auth.New().Me(c)
		c.HTML("dashboard/index", gin.H{
			"title": "Dashboard",
			"admin": admin,
		})
	}
}
