package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Username    string    `gorm:"size:50;not null" json:"username"`
	Email       string    `gorm:"size:100;not null;unique" json:"email"`
	Password    string    `gorm:"size:100;not null" json:"-"`
	Gender      string    `gorm:"size:10" json:"gender"`
	Age         int       `json:"age"`
	Occupation  string    `gorm:"size:100" json:"occupation"`
	Hobbies     string    `gorm:"size:255" json:"hobbies"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	LastLoginAt time.Time `json:"lastLoginAt"`
}

// TableName 设置表名
func (User) TableName() string {
	return "users"
}

// BeforeCreate 创建前的钩子
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// BeforeUpdate 更新前的钩子
func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
