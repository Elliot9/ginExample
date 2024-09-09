package oauth

import (
	"github.com/elliot9/gin-example/internal/pkg/context"
	"github.com/elliot9/gin-example/internal/repository/amqp"
	"github.com/elliot9/gin-example/internal/repository/mysql"
	"github.com/elliot9/gin-example/internal/services/oauth"
	"github.com/go-playground/validator/v10"
)

type Handler interface {
	GetQuery() context.HandlerFunc
	Callback() context.HandlerFunc
}

type handler struct {
	service oauth.Service
}

func New(db mysql.Repo, validator *validator.Validate, amqp amqp.Repo) Handler {
	return &handler{
		service: oauth.New(db, validator, amqp),
	}
}
