package admin

import (
	"github/elliot9/ginExample/internal/pkg/context"
	"net/http"
)

func (h *handler) Logout() context.HandlerFunc {
	return func(c context.Context) {
		c.Session().Delete(context.SessionAuthKey)
		c.Session().Save()

		c.Redirect(http.StatusTemporaryRedirect, "/admin/login")
	}
}
