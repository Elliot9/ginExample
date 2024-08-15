package models

type Admin struct {
	Model    `gorm:"embedded"`
	Name     string `json:"name" gorm:"type:varchar(20)"`
	Email    string `json:"email" gorm:"unique;index;type:varchar(20)"`
	Password string `json:"-" gorm:"type:varchar(255)"`
}
