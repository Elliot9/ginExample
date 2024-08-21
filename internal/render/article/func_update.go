package article

import (
	"github/elliot9/ginExample/internal/auth"
	"github/elliot9/ginExample/internal/pkg/context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type updateArticleReqeust struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Tags    string `json:"tags"`
	Status  bool   `json:"status"`
}

func (h *handler) Update() context.HandlerFunc {
	return func(c context.Context) {
		type uri struct {
			ID int `uri:"id"`
		}
		var uriParam uri
		if err := c.ShouldBindURI(&uriParam); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		if uriParam.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "文章不存在",
			})
			return
		}

		req := new(updateArticleReqeust)
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
		_, err := h.service.FindById(admin, uriParam.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "文章不存在",
			})
			return
		}

		h.service.Update(admin, uriParam.ID, req.Title, req.Content, publishTime, req.Status, req.Tags)

		c.JSON(http.StatusOK, gin.H{
			"message": "文章已更新",
		})
	}
}
