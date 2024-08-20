package article

import (
	"time"
)

func (s *service) Update(id int, title, content string, time *time.Time, status bool, tags ...string) error {
	article, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	article.Title = title
	article.Content = content
	article.Time = time
	article.Status = status

	return s.repo.Update(article)
}
