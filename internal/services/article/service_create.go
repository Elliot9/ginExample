package article

import (
	"github/elliot9/ginExample/internal/models"
	"time"
)

func (s *service) Create(auth *models.Admin, title, content string, time *time.Time, status bool, tags ...string) (int, error) {
	// todo: 处理标签
	return s.repo.Create(&models.Article{
		AdminId: int(auth.ID),
		Title:   title,
		Content: content,
		Time:    time,
		Status:  status,
	})
}
