package router

import (
	"github/elliot9/ginExample/internal/middleware"
	"github/elliot9/ginExample/internal/render/admin"
	"github/elliot9/ginExample/internal/render/dashboard"
)

var renderApi = router(func(r *resource) {
	api := r.mux.Group("/admin")
	{
		admin := admin.New(r.db, r.cache, r.logger, r.validator)
		// guest only
		guestOnlyGroup := api.Group("", middleware.AdaptMiddleware(r.middleware.Guest()))
		guestOnlyGroup.GET("/login", wrapHandler(admin.LoginPage()))
		guestOnlyGroup.GET("/register", wrapHandler(admin.RegisterPage()))
		guestOnlyGroup.POST("/login", wrapHandler(admin.Login()))
		guestOnlyGroup.POST("/register", wrapHandler(admin.Register()))

		// auth only
		authOnlyGroup := api.Group("", middleware.AdaptMiddleware(r.middleware.Auth()))
		authOnlyGroup.GET("/logout", wrapHandler(admin.Logout()))
		authOnlyGroup.GET("", wrapHandler(dashboard.New().IndexPage()))
	}

})
