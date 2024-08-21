package article

import (
	"github/elliot9/ginExample/internal/models"
	"time"
)

func (s *service) Update(auth *models.Admin, id int, title, content string, time *time.Time, status bool, tags ...string) error {
	article, err := s.repo.FindById(int(auth.ID), id)
	if err != nil {
		return err
	}

	article.Title = title
	article.Content = content
	article.Time = time
	article.Status = status

	return s.repo.Update(article)
}
