package models

type User struct {
	Model   `gorm:"embedded"`
	Email   string `json:"email" gorm:"unique;index;type:varchar(20)"`
	Agent   string `json:"agent" gorm:"type:enum('google', 'facebook')"`
	State   bool   `json:"state" gorm:"type:boolean"`
	AgentID string `json:"agent_id" gorm:"type:varchar(255)"`
}
