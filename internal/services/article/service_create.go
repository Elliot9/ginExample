package article

import (
	"context"
	"strconv"
	"time"

	"github.com/elliot9/gin-example/internal/models"
)

func (s *service) Create(auth *models.Admin, title, content string, time *time.Time, status bool, tags ...string) (int, error) {
	id, err := s.repo.Create(&models.Article{
		AdminId: int(auth.ID),
		Title:   title,
		Content: content,
		Time:    time,
		Status:  status,
	})

	if err != nil {
		return 0, err
	}

	cache := s.cache.Tags([]string{"article", "article_" + strconv.Itoa(id)})
	cache.Del(context.Background(), "article_"+strconv.Itoa(id))
	return id, nil
}
