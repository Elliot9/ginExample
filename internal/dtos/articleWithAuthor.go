package dtos

import "github/elliot9/ginExample/internal/models"

type ArticleWithAuthor struct {
	models.Article
	Admin models.Admin `gorm:"foreignKey:AdminId"`
}
