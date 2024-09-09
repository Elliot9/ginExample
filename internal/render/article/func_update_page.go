package article

import (
	"net/http"

	"github.com/elliot9/gin-example/internal/auth"
	"github.com/elliot9/gin-example/internal/pkg/context"

	"github.com/gin-gonic/gin"
)

type updatePageRequest struct {
	ID int `uri:"id"`
}

func (h *handler) UpdatePage() context.HandlerFunc {
	return func(c context.Context) {
		req := new(updatePageRequest)
		if err := c.ShouldBindURI(req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		admin := auth.New().Me(c)

		article, err := h.service.FindById(admin, req.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "文章不存在",
			})
			return
		}

		c.HTML("article/update", gin.H{
			"title":   "更新文章",
			"article": article,
		})
	}
}
