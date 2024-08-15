package models

import "time"

type Model struct {
	ID        uint      `gorm:"not null;primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
