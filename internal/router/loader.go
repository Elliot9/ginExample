package router

import (
	"github/elliot9/ginExample/internal/pkg/context"
	"github/elliot9/ginExample/internal/repository/mysql"
	"github/elliot9/ginExample/internal/repository/redis"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type resource struct {
	mux       *gin.Engine
	db        mysql.Repo
	cache     redis.Repo
	logger    *zap.Logger
	validator *validator.Validate
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
	r := &resource{
		mux:       mux,
		db:        db,
		cache:     cache,
		logger:    zap.New(zapcore.NewNopCore()),
		validator: validator,
	}

	// fetch templates
	r.mux.LoadHTMLGlob("internal/assets/templates/**/*")

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
