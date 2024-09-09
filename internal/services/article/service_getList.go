package article

import (
	"github.com/elliot9/gin-example/internal/models"
	"github.com/elliot9/gin-example/internal/pkg/paginator"
	"github.com/elliot9/gin-example/internal/repository/article"
)

func (s *service) GetList(auth *models.Admin, page int, sortBy string, keyword string) (paginator.Paginator, error) {
	return s.repo.GetList(int(auth.ID), page, pageSize, article.SortBy(sortBy), keyword)
}
