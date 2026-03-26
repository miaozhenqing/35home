package models

import (
	"time"
)

// Vote 投票模型
type Vote struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"userId"`
	Status    string    `json:"status"` // employed 或 unemployed
	City      string    `json:"city"`
	Industry  string    `json:"industry"`
	Occupation string    `json:"occupation"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"createdAt"`
}

// TableName 指定表名
func (Vote) TableName() string {
	return "votes"
}
