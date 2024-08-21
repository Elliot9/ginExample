package router

import (
	"github/elliot9/ginExample/internal/api/health"
	"github/elliot9/ginExample/internal/middleware"
	"github/elliot9/ginExample/internal/render/article"
)

var apiRouter = router(func(r *resource) {
	api := r.mux.Group("/api")
	{
		// 健康檢查
		api.GET("/health", wrapHandler(health.New().Ping()))

		// auth only
		apiAuthGroup := api.Group("", middleware.AdaptMiddleware(r.middleware.Auth()))
		{
			article := article.New(r.db, r.cache, r.validator)
			apiAuthGroup.POST("/articles/create", wrapHandler(article.Create()))
			apiAuthGroup.POST("/articles/temporary", wrapHandler(article.Temporary()))
			apiAuthGroup.POST("/articles/:id/update", wrapHandler(article.Update()))
			apiAuthGroup.POST("/articles/:id/delete", wrapHandler(article.Delete()))
		}
	}
})
