package article

import (
	"context"
	"strconv"

	"github.com/elliot9/gin-example/internal/models"
)

func (s *service) Delete(auth *models.Admin, id int) error {
	article, err := s.repo.FindByID(int(auth.ID), id)
	if err != nil {
		return err
	}

	cache := s.cache.Tags([]string{"article", "article_" + strconv.Itoa(id)})
	cache.Del(context.Background(), "article_"+strconv.Itoa(id))
	return s.repo.Delete(article)
}
