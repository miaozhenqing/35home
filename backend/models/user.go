package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Username    string `gorm:"size:50;not null;unique" json:"username"`
	Email       string `gorm:"size:100;not null;unique" json:"email"`
	Password    string `gorm:"size:100;not null" json:"-"`
	Gender      string `gorm:"size:10" json:"gender"`
	Age         int    `json:"age"`
	Occupation  string `gorm:"size:100" json:"occupation"`
	Hobbies     string `gorm:"size:255" json:"hobbies"`
	CreatedAt   int64  `gorm:"column:created_at;type:bigint" json:"createdAt"`
	UpdatedAt   int64  `gorm:"column:updated_at;type:bigint" json:"updatedAt"`
	LastLoginAt int64  `gorm:"column:last_login_at;type:bigint" json:"lastLoginAt"`
}

// TableName 设置表名
func (User) TableName() string {
	return "users"
}

// BeforeCreate 创建前的钩子
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	now := time.Now().UnixMilli()
	scope.SetColumn("CreatedAt", now)
	scope.SetColumn("UpdatedAt", now)
	return nil
}

// BeforeUpdate 更新前的钩子
func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().UnixMilli())
	return nil
}
