package user

import (
	"github.com/elliot9/gin-example/internal/models"
	"github.com/elliot9/gin-example/internal/repository/mysql"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(newUser *models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	UpdateRefreshToken(userRefreshToken *models.UserRefreshToken) error
}

type userRepo struct {
	db mysql.Repo
}

func New(db mysql.Repo) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (repo *userRepo) Create(newUser *models.User) (*models.User, error) {
	result := repo.db.GetDbW().Create(newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return newUser, nil
}

func (repo *userRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := repo.db.GetDbR().Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *userRepo) UpdateRefreshToken(userRefreshToken *models.UserRefreshToken) error {
	// if exists then update, else create
	var user models.UserRefreshToken
	result := repo.db.GetDbW().Where("user_id = ?", userRefreshToken.UserID).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	if result.Error == gorm.ErrRecordNotFound {
		result = repo.db.GetDbW().Create(userRefreshToken)
	} else {
		result = repo.db.GetDbW().Model(&user).UpdateColumns(userRefreshToken)
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}
