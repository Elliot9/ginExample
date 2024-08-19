package admin

import (
	"github/elliot9/ginExample/internal/pkg/context"

	"github.com/gin-gonic/gin"
)

func (h *handler) LoginPage() context.HandlerFunc {
	return func(c context.Context) {
		var messages gin.H

		flashMessages, ok := c.GetFlash()
		if ok && len(flashMessages) > 0 {
			messages = gin.H{}
			for key, value := range flashMessages {
				messages[key] = value
			}
		}

		c.HTML("login", gin.H{
			"messages": messages,
		})
	}
}
