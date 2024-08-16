package dashboard

import "github/elliot9/ginExample/internal/pkg/context"

func (h *handler) IndexPage() context.HandlerFunc {
	return func(c context.Context) {
		c.HTML("dashboard/index.html", nil)
	}
}
