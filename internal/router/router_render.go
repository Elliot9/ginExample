package router

import (
	"github/elliot9/ginExample/internal/render/admin"
)

var renderApi = router(func(r *resource) {
	api := r.mux.Group("/admin")
	{
		admin := admin.New()
		// 登入頁面
		api.GET("/login", wrapHandler(admin.LoginPage()))
		// 註冊頁面
		api.GET("/register", wrapHandler(admin.RegisterPage()))
	}
})
