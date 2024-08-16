package admin

import (
	"errors"
	"github/elliot9/ginExample/internal/models"
)

func (s *service) Login(email, password string) (*models.Admin, error) {
	admin := s.repo.FindByEmail(email)
	if admin == nil {
		return admin, errors.New("帳號不存在")
	}

	if result, err := s.hash.Verify(admin.Password, password); !result || err != nil {
		return admin, errors.New("密碼錯誤")
	}

	return admin, nil
}
