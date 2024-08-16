package admin

import (
	"github/elliot9/ginExample/internal/models"
	"github/elliot9/ginExample/internal/pkg/hash"
	"github/elliot9/ginExample/internal/repository/admin"
	"github/elliot9/ginExample/internal/repository/mysql"

	"github.com/go-playground/validator/v10"
)

var _ Service = (*service)(nil)

type Service interface {
	Register(name, email, password string) (int, error)
	Login(email, password string) (*models.Admin, error)
}

type service struct {
	validator *validator.Validate
	repo      admin.AdminRepo
	hash      hash.Hash
}

func New(db mysql.Repo, validator *validator.Validate) Service {
	return &service{
		validator: validator,
		repo:      admin.New(db),
		hash:      hash.New(),
	}
}
