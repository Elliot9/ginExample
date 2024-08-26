package article

import (
	"github/elliot9/ginExample/internal/pkg/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Get() context.HandlerFunc {
	return func(c context.Context) {
		type uri struct {
			ID int `uri:"id"`
		}
		var uriParam uri
		if err := c.ShouldBindURI(&uriParam); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		article, err := h.service.GetDetailByID(uriParam.ID, true)
		if article == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "找不到文章",
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusOK, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "讀取成功",
			"title":      article.Title,
			"content":    article.Content,
			"created_at": article.CreatedAt,
			"author":     article.Admin.Name,
		})
		return
	}
}
