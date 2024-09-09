package article

import (
	"time"

	"github.com/elliot9/gin-example/internal/dtos"
	"github.com/elliot9/gin-example/internal/models"
	"github.com/elliot9/gin-example/internal/pkg/cache"
	"github.com/elliot9/gin-example/internal/pkg/paginator"
	"github.com/elliot9/gin-example/internal/repository/article"
	"github.com/elliot9/gin-example/internal/repository/mysql"
	"github.com/elliot9/gin-example/internal/repository/redis"

	"github.com/go-playground/validator/v10"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(auth *models.Admin, title, content string, time *time.Time, status bool, tags ...string) (int, error)
	Update(auth *models.Admin, id int, title, content string, time *time.Time, status bool, tags ...string) error
	GetList(auth *models.Admin, page int, sortBy string, keyword string) (paginator.Paginator, error)
	FindById(auth *models.Admin, id int) (*models.Article, error)
	Delete(auth *models.Admin, id int) error

	// api
	GetAllList(page int, onlyActive bool) (paginator.Paginator, error)
	GetDetailByID(id int, onlyActive bool) (*dtos.ArticleWithAuthor, error)
}

const pageSize = 10

type service struct {
	validator *validator.Validate
	repo      article.ArticleRepo
	cache     cache.Cache
}

func New(db mysql.Repo, cacheRepo redis.Repo, validator *validator.Validate) Service {
	return &service{
		validator: validator,
		repo:      article.New(db),
		cache:     cache.New(cacheRepo),
	}
}
