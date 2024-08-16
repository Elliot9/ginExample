package admin

import "github/elliot9/ginExample/internal/pkg/context"

func (h *handler) RegisterPage() context.HandlerFunc {
	return func(c context.Context) {
		c.HTML("admin/register.html", nil)
	}
}
