package article

import (
	"github/elliot9/ginExample/internal/models"
	"github/elliot9/ginExample/internal/pkg/paginator"
	"github/elliot9/ginExample/internal/repository/article"
	"github/elliot9/ginExample/internal/repository/mysql"
	"time"

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
}

const pageSize = 10

type service struct {
	validator *validator.Validate
	repo      article.ArticleRepo
}

func New(db mysql.Repo, validator *validator.Validate) Service {
	return &service{
		validator: validator,
		repo:      article.New(db),
	}
}
