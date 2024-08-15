package admin

import (
	"github/elliot9/ginExample/internal/models"
)

func (s *service) Register(name, email, password string) (id int, err error) {
	password, _ = s.hash.Hash(password)
	return s.repo.Create(&models.Admin{
		Name:     name,
		Email:    email,
		Password: password,
	})
}
