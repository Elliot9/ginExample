package admin

import (
	"errors"
	"github/elliot9/ginExample/internal/models"
	"github/elliot9/ginExample/internal/repository/mysql"
)

type AdminRepo interface {
	Create(newAdmin *models.Admin) (id int, err error)
	Exists(email string) bool
	FindByEmail(email string) *models.Admin
}

type adminRepo struct {
	db mysql.Repo
}

func New(db mysql.Repo) AdminRepo {
	return &adminRepo{
		db: db,
	}
}

func (repo *adminRepo) Create(newAdmin *models.Admin) (id int, err error) {
	if repo.Exists(newAdmin.Email) {
		return 0, errors.New("帳號已存在")
	}

	result := repo.db.GetDbW().Create(newAdmin)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(newAdmin.ID), nil
}

func (repo *adminRepo) Exists(email string) bool {
	return repo.FindByEmail(email) != nil
}

func (repo *adminRepo) FindByEmail(email string) *models.Admin {
	var admin models.Admin
	result := repo.db.GetDbR().Where("email = ?", email).First(&admin)

	if result.Error != nil {
		return nil
	}

	return &admin
}
