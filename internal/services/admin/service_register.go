package admin

import (
	"github.com/elliot9/gin-example/internal/models"
)

func (s *service) Register(name, email, password string) (id int, err error) {
	password, _ = s.hash.Hash(password)
	return s.repo.Create(&models.Admin{
		Name:     name,
		Email:    email,
		Password: password,
	})
}
