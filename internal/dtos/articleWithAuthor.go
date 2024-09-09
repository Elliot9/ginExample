package dtos

import (
	"github.com/elliot9/gin-example/internal/models"
)

type ArticleWithAuthor struct {
	models.Article
	Admin models.Admin `gorm:"foreignKey:AdminId"`
}
