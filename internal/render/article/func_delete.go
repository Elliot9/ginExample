package article

import (
	"net/http"

	"github.com/elliot9/gin-example/internal/auth"
	"github.com/elliot9/gin-example/internal/pkg/context"

	"github.com/gin-gonic/gin"
)

func (h *handler) Delete() context.HandlerFunc {
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

		admin := auth.New().Me(c)
		_, err := h.service.FindById(admin, uriParam.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "文章不存在",
			})
			return
		}

		h.service.Delete(admin, uriParam.ID)

		c.JSON(http.StatusOK, gin.H{
			"message": "文章已更刪除",
		})
	}
}
