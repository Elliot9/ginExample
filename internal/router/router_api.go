package router

import (
	"github.com/elliot9/gin-example/internal/api/health"
	"github.com/elliot9/gin-example/internal/api/oauth"
	"github.com/elliot9/gin-example/internal/middleware"
	"github.com/elliot9/gin-example/internal/render/article"
)

var apiRouter = router(func(r *resource) {
	api := r.mux.Group("/api")
	{
		article := article.New(r.db, r.cache, r.validator)
		oauth := oauth.New(r.db, r.validator, r.amqp)

		// 健康檢查
		api.GET("/health", wrapHandler(health.New().Ping()))

		// oauth
		oauthGroup := api.Group("/auth")
		{
			oauthGroup.GET("/:agent", wrapHandler(oauth.GetQuery()))
			oauthGroup.POST("/:agent/callback", wrapHandler(oauth.Callback()))
		}

		// admin only
		apiAuthGroup := api.Group("/admin", middleware.AdaptMiddleware(r.middleware.Auth()))
		{
			apiAuthGroup.POST("/articles/create", wrapHandler(article.Create()))
			apiAuthGroup.POST("/articles/temporary", wrapHandler(article.Temporary()))
			apiAuthGroup.POST("/articles/:id/update", wrapHandler(article.Update()))
			apiAuthGroup.POST("/articles/:id/delete", wrapHandler(article.Delete()))
		}

		// user only
		// todo jwt login user
		apiUserGroup := api.Group("/articles")
		{
			apiUserGroup.GET("", wrapHandler(article.List()))
			apiUserGroup.GET("/:id", wrapHandler(article.Get()))
		}
	}
})
