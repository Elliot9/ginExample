package article

import (
	"github/elliot9/ginExample/internal/auth"
	"github/elliot9/ginExample/internal/pkg/context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ListQuery struct {
	Page    int    `form:"page,default=1"`
	SortBy  string `form:"sort,default=createdAt"`
	Keyword string `form:"keyword"`
}

func (h *handler) ListPage() context.HandlerFunc {
	return func(c context.Context) {
		req := new(ListQuery)
		if err := c.ShouldBindQuery(req); err != nil {
			errors := make(map[string]any)
			for _, fieldErr := range err.(validator.ValidationErrors) {
				errors[fieldErr.Field()] = fieldErr.ActualTag()
			}

			c.JSON(http.StatusBadRequest, errors)
			return
		}

		admin := auth.New().Me(c)
		pg, err := h.service.GetList(admin, req.Page, req.SortBy, req.Keyword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "文章失敗",
			})
			return
		}

		totalPagesMap := make(map[int]any, pg.TotalPages())
		for i := 1; i <= pg.TotalPages(); i++ {
			totalPagesMap[i] = i
		}

		c.HTML("article/list", gin.H{
			"title":         "文章列表",
			"admin":         admin,
			"pg":            pg,
			"SortBy":        req.SortBy,
			"Keyword":       req.Keyword,
			"totalPagesMap": totalPagesMap,
		})
	}
}
