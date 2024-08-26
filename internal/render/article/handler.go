package article

import (
	"github/elliot9/ginExample/internal/pkg/context"
	"github/elliot9/ginExample/internal/repository/mysql"
	"github/elliot9/ginExample/internal/repository/redis"
	"github/elliot9/ginExample/internal/services/article"

	"github.com/go-playground/validator/v10"
)

type Handler interface {
	CreatePage() context.HandlerFunc
	Create() context.HandlerFunc
	ListPage() context.HandlerFunc
	Temporary() context.HandlerFunc
	UpdatePage() context.HandlerFunc
	Update() context.HandlerFunc
	Delete() context.HandlerFunc
	List() context.HandlerFunc
	Get() context.HandlerFunc
}

type handler struct {
	service article.Service
}

func New(db mysql.Repo, cache redis.Repo, validator *validator.Validate) Handler {
	return &handler{
		service: article.New(db, cache, validator),
	}
}
