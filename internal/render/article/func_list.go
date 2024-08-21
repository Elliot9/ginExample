package article

import (
	"encoding/json"
	"github/elliot9/ginExample/internal/dtos"
	"github/elliot9/ginExample/internal/pkg/context"
	"net/http"
	"time"

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
		page := 1

		if body := c.Body(); len(body) != 0 {
			var data map[string]interface{}
			if err := json.Unmarshal(body, &data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			page = int(data["page"].(float64))
		}

		pg, err := h.service.GetAllList(page, false)
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
