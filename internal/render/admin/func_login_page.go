package admin

import "github/elliot9/ginExample/internal/pkg/context"

func (h *handler) LoginPage() context.HandlerFunc {
	return func(c context.Context) {
		c.HTML("admin/login.html", nil)
	}
}
