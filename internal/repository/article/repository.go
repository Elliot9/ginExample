package article

import (
	"github/elliot9/ginExample/internal/models"
	"github/elliot9/ginExample/internal/pkg/paginator"
	"github/elliot9/ginExample/internal/repository/mysql"
)

type SortBy string

const (
	SortByCreatedAt SortBy = "createdAt"
	SortByStatus    SortBy = "status"
	SortByTitle     SortBy = "title"
)

type ArticleRepo interface {
	Create(newArticle *models.Article) (id int, err error)
	FindById(id int) (*models.Article, error)
	Update(article *models.Article) error
	GetList(page, pageSize int, sortBy SortBy, keyword string) (paginator.Paginator, error)
}

type articleRepo struct {
	db mysql.Repo
}

func New(db mysql.Repo) ArticleRepo {
	return &articleRepo{
		db: db,
	}
}

func (repo *articleRepo) Create(newArticle *models.Article) (id int, err error) {
	result := repo.db.GetDbW().Create(newArticle)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(newArticle.ID), nil
}

func (repo *articleRepo) FindById(id int) (*models.Article, error) {
	var article models.Article
	result := repo.db.GetDbR().First(&article, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &article, nil
}

func (repo *articleRepo) Update(article *models.Article) error {
	result := repo.db.GetDbW().Save(article)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *articleRepo) GetList(page, pageSize int, sortBy SortBy, keyword string) (paginator.Paginator, error) {
	query := repo.db.GetDbR().Model(&models.Article{})

	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}

	switch sortBy {
	case SortByCreatedAt:
		query = query.Order("created_at DESC")
	case SortByStatus:
		query = query.Order("status DESC")
	case SortByTitle:
		query = query.Order("title ASC")
	}

	return paginator.NewPaginator(query, page, pageSize, &[]models.Article{})
}
