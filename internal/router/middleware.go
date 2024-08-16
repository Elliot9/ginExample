package router

import (
	"github/elliot9/ginExample/internal/middleware"
)

type RouterMiddleware interface {
	CheckLogin() middleware.Middleware
	CheckPermission() middleware.Middleware
	Guest() middleware.Middleware
	Auth() middleware.Middleware
}

type routerMiddleware struct{}

func (r *routerMiddleware) CheckLogin() middleware.Middleware {
	return middleware.NewAuthMiddleware()
}

func (r *routerMiddleware) CheckPermission() middleware.Middleware {
	return middleware.NewAuthMiddleware()
}

func (r *routerMiddleware) Guest() middleware.Middleware {
	return middleware.NewGuestMiddleware()
}

func (r *routerMiddleware) Auth() middleware.Middleware {
	return middleware.NewAuthMiddleware()
}

func newRouterMiddleware() RouterMiddleware {
	return &routerMiddleware{}
}
