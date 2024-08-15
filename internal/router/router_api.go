package router

import (
	"github/elliot9/ginExample/internal/api/admin"
	"github/elliot9/ginExample/internal/api/health"
)

var apiRouter = router(func(r *resource) {
	api := r.mux.Group("/api")
	{
		// 健康檢查
		api.GET("/health", wrapHandler(health.New().Ping()))

		admin := admin.New(r.db, r.cache, r.logger, r.validator)
		// todo: 增加 middleware
		api.POST("/admin/login", wrapHandler(admin.Login()))
		api.POST("/admin/register", wrapHandler(admin.Register()))
	}
})
