package article

import (
	"github/elliot9/ginExample/internal/pkg/paginator"
	"github/elliot9/ginExample/internal/repository/article"
)

func (s *service) GetAllList(page int, onlyActive bool) (paginator.Paginator, error) {
	return s.repo.GetAllList(page, pageSize, article.SortByCreatedAt, onlyActive)
}
