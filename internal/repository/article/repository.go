package article

import (
	"github.com/elliot9/gin-example/internal/dtos"
	"github.com/elliot9/gin-example/internal/models"
	"github.com/elliot9/gin-example/internal/pkg/paginator"
	"github.com/elliot9/gin-example/internal/repository/mysql"
)

type SortBy string

const (
	SortByCreatedAt SortBy = "createdAt"
	SortByStatus    SortBy = "status"
	SortByTitle     SortBy = "title"
)

type ArticleRepo interface {
	Create(newArticle *models.Article) (id int, err error)
	FindByID(adminId int, id int) (*models.Article, error)
	Update(article *models.Article) error
	GetList(adminId int, page, pageSize int, sortBy SortBy, keyword string) (paginator.Paginator, error)
	Delete(article *models.Article) error
	GetAllList(page, pageSize int, sortBy SortBy, onlyActive bool) (paginator.Paginator, error)
	GetDetailByID(id int, onlyActive bool) (*dtos.ArticleWithAuthor, error)
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

func (repo *articleRepo) FindByID(adminId int, id int) (*models.Article, error) {
	var article models.Article
	result := repo.db.GetDbR().Where("admin_id = ?", adminId).First(&article, id)
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

func (repo *articleRepo) Delete(article *models.Article) error {
	result := repo.db.GetDbW().Delete(article)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *articleRepo) GetList(adminId int, page, pageSize int, sortBy SortBy, keyword string) (paginator.Paginator, error) {
	query := repo.db.GetDbR().Model(&models.Article{}).Where("admin_id = ?", adminId)

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

func (repo *articleRepo) GetAllList(page, pageSize int, sortBy SortBy, onlyActive bool) (paginator.Paginator, error) {
	query := repo.db.GetDbR().Model(&models.Article{})

	switch sortBy {
	case SortByCreatedAt:
		query = query.Order("created_at DESC")
	case SortByStatus:
		query = query.Order("status DESC")
	case SortByTitle:
		query = query.Order("title ASC")
	}

	if onlyActive {
		query = query.Where("status = ?", true)
	}

	return paginator.NewPaginatorWithPreload(query, page, pageSize, &[]dtos.ArticleWithAuthor{}, "Admin")
}

func (repo *articleRepo) GetDetailByID(id int, onlyActive bool) (*dtos.ArticleWithAuthor, error) {
	var article dtos.ArticleWithAuthor
	query := repo.db.GetDbR().Model(&models.Article{}).Where("id = ?", id).Preload("Admin")

	if onlyActive {
		query = query.Where("status = ?", true)
	}

	result := query.First(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return &article, nil
}
