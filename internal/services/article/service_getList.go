package article

import (
	"github/elliot9/ginExample/internal/pkg/paginator"
	"github/elliot9/ginExample/internal/repository/article"
)

const pageSize = 10

func (s *service) GetList(page int, sortBy string, keyword string) (paginator.Paginator, error) {
	return s.repo.GetList(page, pageSize, article.SortBy(sortBy), keyword)
}
