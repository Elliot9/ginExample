package router

import (
	"github/elliot9/ginExample/internal/middleware"
	"github/elliot9/ginExample/internal/pkg/context"
	"github/elliot9/ginExample/internal/repository/mysql"
	"github/elliot9/ginExample/internal/repository/redis"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type resource struct {
	mux        *gin.Engine
	db         mysql.Repo
	cache      redis.Repo
	logger     *zap.Logger
	validator  *validator.Validate
	middleware RouterMiddleware
}

var _ routerRegister = (router)(nil)

type routerRegister interface {
	register(r *resource)
}

type router func(r *resource)

func (ro router) register(r *resource) {
	ro(r)
}

func RegisterRouter(mux *gin.Engine, db mysql.Repo, cache redis.Repo, validator *validator.Validate) {
	logger := zap.NewExample()
	r := &resource{
		mux:        mux,
		db:         db,
		cache:      cache,
		logger:     logger,
		validator:  validator,
		middleware: newRouterMiddleware(),
	}

	r.mux.StaticFS("/assets", http.Dir("internal/assets"))
	r.mux.LoadHTMLGlob("internal/assets/templates/**/*")

	// global middlewares
	r.mux.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))
	r.mux.Use(middleware.AdaptMiddleware(middleware.NewErrorMiddleware()))
	r.mux.Use(middleware.AdaptMiddleware(middleware.NewLoggerMiddleware(logger)))

	for _, re := range getAllRouterRegister() {
		re.register(r)
	}
}

func getAllRouterRegister() []routerRegister {
	return []routerRegister{apiRouter, renderApi}
}

func wrapHandler(handler context.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(context.NewContext(c))
	}
}
