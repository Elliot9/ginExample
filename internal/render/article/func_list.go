package article

import (
	"net/http"
	"time"

	"github.com/elliot9/gin-example/internal/dtos"
	"github.com/elliot9/gin-example/internal/pkg/context"

	"github.com/gin-gonic/gin"
)

type articleItem struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Author    string    `json:"author"`
}

func (h *handler) List() context.HandlerFunc {
	return func(c context.Context) {
		var pageQuery struct {
			Page int `form:"page,default=1"`
		}

		c.ShouldBindQuery(&pageQuery)

		pg, err := h.service.GetAllList(pageQuery.Page, false)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "讀取失敗",
			})
			return
		}

		pgItems := pg.GetItems().(*[]dtos.ArticleWithAuthor)
		items := make([]articleItem, len(*pgItems))

		for i, item := range *pgItems {
			items[i] = articleItem{
				ID:        item.ID,
				Title:     item.Title,
				CreatedAt: item.CreatedAt,
				Author:    item.Admin.Name,
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "讀取成功",
			"items":      items,
			"total":      pg.Total(),
			"page":       pg.Page(),
			"totalPages": pg.TotalPages(),
		})
	}
}
