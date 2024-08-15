package admin

import "github/elliot9/ginExample/internal/pkg/context"

type Handler interface {
	LoginPage() context.HandlerFunc
	RegisterPage() context.HandlerFunc
}

type handler struct {
}

func (h *handler) LoginPage() context.HandlerFunc {
	return func(c context.Context) {
		c.HTML("admin/login.html", nil)
	}
}

func (h *handler) RegisterPage() context.HandlerFunc {
	return func(c context.Context) {
		c.HTML("admin/register.html", nil)
	}
}

func New() Handler {
	return &handler{}
}
