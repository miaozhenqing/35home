package config

import (
	"fmt"
	"log"

	"35home/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDatabase 初始化数据库连接 todo 通过外部配置，设置数据库连接参数
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

	// 自动迁移表结构 todo 只在开发环境中使用，生产环境中不应该自动迁移
	err = DB.AutoMigrate(&models.User{}, &models.Vote{}).Error
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
