package config

import (
	"fmt"
	"log"

	"35home/models"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase(config *Config) error {
	var err error
	
	// 连接数据库
	DB, err = gorm.Open("mysql", config.GetDSN())
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	
	// 设置连接池
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	
	// 自动迁移表结构
	err = DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	
	log.Println("Database connected and migrated successfully")
	return nil
}

// CloseDatabase 关闭数据库连接
func CloseDatabase() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
