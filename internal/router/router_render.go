package router

import (
	"github/elliot9/ginExample/internal/render/admin"
)

var renderApi = router(func(r *resource) {
	api := r.mux.Group("/admin")
	{
		// 登入頁面
		api.GET("/login", wrapHandler(admin.New().LoginPage()))
	}
})
