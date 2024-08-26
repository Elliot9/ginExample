package article

import (
	"context"
	"github/elliot9/ginExample/internal/models"
	"strconv"
	"time"
)

func (s *service) Update(auth *models.Admin, id int, title, content string, time *time.Time, status bool, tags ...string) error {
	article, err := s.repo.FindByID(int(auth.ID), id)
	if err != nil {
		return err
	}

	article.Title = title
	article.Content = content
	article.Time = time
	article.Status = status

	err = s.repo.Update(article)
	if err != nil {
		return err
	}

	cache := s.cache.Tags([]string{"article", "article_" + strconv.Itoa(id)})
	cache.Del(context.Background(), "article_"+strconv.Itoa(id))
	return nil
}
