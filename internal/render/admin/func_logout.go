package admin

import (
	"net/http"

	"github.com/elliot9/gin-example/internal/pkg/context"
)

func (h *handler) Logout() context.HandlerFunc {
	return func(c context.Context) {
		c.Session().Delete(context.SessionAuthKey)
		c.Session().Save()

		c.Redirect(http.StatusTemporaryRedirect, "/admin/login")
	}
}
