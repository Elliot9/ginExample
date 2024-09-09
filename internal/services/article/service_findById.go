package article

import (
	"github.com/elliot9/gin-example/internal/models"
)

func (s *service) FindById(auth *models.Admin, id int) (*models.Article, error) {
	return s.repo.FindByID(int(auth.ID), id)
}
