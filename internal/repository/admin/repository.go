package admin

import (
	"errors"
	"github/elliot9/ginExample/internal/models"
	"github/elliot9/ginExample/internal/repository/mysql"

	"gorm.io/gorm"
)

type AdminRepo interface {
	Create(newAdmin *models.Admin) (id int, err error)
	Exists(email string) bool
	//Find(email, password string) (*models.Admin, error)
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
	var admin models.Admin

	result := repo.db.GetDbR().Where("email = ?", email).First(&admin)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}
