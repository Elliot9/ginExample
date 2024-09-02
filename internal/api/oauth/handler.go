package oauth

import (
	"github/elliot9/ginExample/internal/pkg/context"
	"github/elliot9/ginExample/internal/repository/amqp"
	"github/elliot9/ginExample/internal/repository/mysql"
	"github/elliot9/ginExample/internal/services/oauth"

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
