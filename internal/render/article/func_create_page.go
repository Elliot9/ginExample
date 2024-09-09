package article

import (
	"github.com/elliot9/gin-example/internal/auth"
	"github.com/elliot9/gin-example/internal/pkg/context"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreatePage() context.HandlerFunc {
	return func(c context.Context) {
		admin := auth.New().Me(c)
		c.HTML("article/create", gin.H{
			"title": "新增文章",
			"admin": admin,
		})
	}
}
