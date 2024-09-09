package article

import (
	"github.com/elliot9/gin-example/internal/pkg/paginator"
	"github.com/elliot9/gin-example/internal/repository/article"
)

func (s *service) GetAllList(page int, onlyActive bool) (paginator.Paginator, error) {
	return s.repo.GetAllList(page, pageSize, article.SortByCreatedAt, onlyActive)
}
