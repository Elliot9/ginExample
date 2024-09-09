package article

import (
	"net/http"
	"time"

	"github.com/elliot9/gin-example/internal/auth"
	"github.com/elliot9/gin-example/internal/pkg/context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type temporaryArticleReqeust struct {
	Id      int    `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Tags    string `json:"tags"`
}

func (h *handler) Temporary() context.HandlerFunc {
	return func(c context.Context) {
		req := new(temporaryArticleReqeust)
		if err := c.ShouldBindJson(req); err != nil {
			errors := make(map[string]any)
			for _, fieldErr := range err.(validator.ValidationErrors) {
				errors[fieldErr.Field()] = fieldErr.ActualTag()
			}

			c.JSON(http.StatusBadRequest, errors)
			return
		}

		var publishTime *time.Time
		if req.Time != "" {
			loc, _ := time.LoadLocation("Asia/Taipei")
			parsedTime, err := time.ParseInLocation("2006-01-02 15:04", req.Time, loc)
			publishTime = &parsedTime
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "無效的時間格式",
				})
				return
			}
		}
		admin := auth.New().Me(c)
		if req.Id == 0 {
			id, err := h.service.Create(admin, req.Title, req.Content, publishTime, false, req.Tags)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "文章暫存失敗",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "文章已暫存",
				"id":      id,
			})
			return
		}

		h.service.Update(admin, req.Id, req.Title, req.Content, publishTime, false, req.Tags)
	}
}
