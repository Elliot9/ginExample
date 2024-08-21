package models

import "time"

type Article struct {
	Model   `gorm:"embedded"`
	Title   string     `json:"title" gorm:"type:varchar(20)"`
	Content string     `json:"content" gorm:"type:longtext"`
	Status  bool       `json:"status" gorm:"type:bool;default:false;index"`
	Time    *time.Time `json:"time,omitempty" gorm:"type:datetime;nullable"`
	AdminId int        `json:"adminId" gorm:"type:uint;index"`
}
