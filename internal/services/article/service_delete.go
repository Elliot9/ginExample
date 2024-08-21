package article

import (
	"github/elliot9/ginExample/internal/models"
)

func (s *service) Delete(auth *models.Admin, id int) error {
	article, err := s.repo.FindById(int(auth.ID), id)
	if err != nil {
		return err
	}

	return s.repo.Delete(article)
}
