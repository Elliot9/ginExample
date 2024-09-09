package article

import (
	"net/http"
	"time"

	"github.com/elliot9/gin-example/internal/auth"
	"github.com/elliot9/gin-example/internal/pkg/context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type publishArticleReqeust struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content"`
	Time    string `form:"time"`
	Tags    string `form:"tags"`
}

func (h *handler) Create() context.HandlerFunc {
	return func(c context.Context) {
		req := new(publishArticleReqeust)
		if err := c.ShouldBindForm(req); err != nil {
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
		id, err := h.service.Create(admin, req.Title, req.Content, publishTime, true, req.Tags)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "文章發佈失敗",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "文章已發佈",
			"id":      id,
		})
	}
}
