package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Vote 投票模型
type Vote struct {
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID     uint64 `json:"userId"`
	Status     string `json:"status"` // employed 或 unemployed
	City       string `json:"city"`
	Industry   string `json:"industry"`
	Occupation string `json:"occupation"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	CreatedAt  int64  `gorm:"type:bigint" json:"createdAt"`
}

// TableName 指定表名
func (Vote) TableName() string {
	return "votes"
}

// BeforeCreate 创建前的钩子
func (v *Vote) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().UnixMilli())
	return nil
}
