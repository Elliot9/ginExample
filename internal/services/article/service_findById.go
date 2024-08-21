package article

import (
	"github/elliot9/ginExample/internal/models"
)

func (s *service) FindById(auth *models.Admin, id int) (*models.Article, error) {
	return s.repo.FindById(int(auth.ID), id)
}
