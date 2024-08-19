package dashboard

import (
	"github/elliot9/ginExample/internal/auth"
	"github/elliot9/ginExample/internal/pkg/context"

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
