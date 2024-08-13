package router

import (
	"github/elliot9/ginExample/internal/api/health"
)

var apiRouter = router(func(r *resource) {
	api := r.mux.Group("/api")
	{
		// 健康檢查
		api.GET("/health", wrapHandler(health.New().Ping()))
	}
})
