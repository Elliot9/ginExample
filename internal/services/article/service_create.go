package article

import (
	"github/elliot9/ginExample/internal/models"
	"time"
)

func (s *service) Create(title, content string, time *time.Time, status bool, tags ...string) (int, error) {
	// todo tags
	return s.repo.Create(&models.Article{
		Title:   title,
		Content: content,
		Time:    time,
		Status:  status,
	})
}
