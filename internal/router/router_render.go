package router

import (
	"github.com/elliot9/gin-example/internal/middleware"
	"github.com/elliot9/gin-example/internal/render/admin"
	"github.com/elliot9/gin-example/internal/render/article"
	"github.com/elliot9/gin-example/internal/render/dashboard"
)

var renderApi = router(func(r *resource) {
	render := r.mux.Group("/admin")
	{
		admin := admin.New(r.db, r.cache, r.logger, r.validator)
		// guest only
		guestOnlyGroup := render.Group("", middleware.AdaptMiddleware(r.middleware.Guest()))
		guestOnlyGroup.GET("/login", wrapHandler(admin.LoginPage()))
		guestOnlyGroup.GET("/register", wrapHandler(admin.RegisterPage()))
		guestOnlyGroup.POST("/login", wrapHandler(admin.Login()))
		guestOnlyGroup.POST("/register", wrapHandler(admin.Register()))

		// auth only
		authOnlyGroup := render.Group("", middleware.AdaptMiddleware(r.middleware.Auth()))
		authOnlyGroup.GET("/logout", wrapHandler(admin.Logout()))
		authOnlyGroup.GET("", wrapHandler(dashboard.New().IndexPage()))

		// article
		article := article.New(r.db, r.cache, r.validator)
		authOnlyGroup.GET("/articles/create", wrapHandler(article.CreatePage()))
		authOnlyGroup.GET("/articles", wrapHandler(article.ListPage()))
		authOnlyGroup.GET("/articles/:id/update", wrapHandler(article.UpdatePage()))
	}
})
