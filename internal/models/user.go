package models

import "time"

type User struct {
	Model   `gorm:"embedded"`
	Name    string `json:"name" gorm:"type:varchar(20)"`
	Email   string `json:"email" gorm:"unique;index;type:varchar(20)"`
	Agent   string `json:"agent" gorm:"type:enum('google', 'facebook')"`
	State   bool   `json:"state" gorm:"type:boolean"`
	AgentID string `json:"agent_id" gorm:"type:varchar(255)"`
}

type UserRefreshToken struct {
	Model     `gorm:"embedded"`
	UserID    uint      `json:"user_id" gorm:"unique"`
	Token     string    `json:"token" gorm:"type:varchar(255)"`
	ExpiredAt time.Time `json:"expired_at" gorm:"type:datetime"`
}
