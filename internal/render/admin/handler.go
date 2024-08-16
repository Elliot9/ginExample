package admin

import (
	"github/elliot9/ginExample/internal/pkg/context"
	"github/elliot9/ginExample/internal/repository/mysql"
	"github/elliot9/ginExample/internal/repository/redis"
	"github/elliot9/ginExample/internal/services/admin"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Handler interface {
	LoginPage() context.HandlerFunc
	RegisterPage() context.HandlerFunc
	Register() context.HandlerFunc
	Login() context.HandlerFunc
	Logout() context.HandlerFunc
}

type handler struct {
	cache   redis.Repo
	logger  *zap.Logger
	service admin.Service
}

func New(db mysql.Repo, cache redis.Repo, logger *zap.Logger, validator *validator.Validate) Handler {
	return &handler{
		cache:   cache,
		logger:  logger,
		service: admin.New(db, validator),
	}
}
