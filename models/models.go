package models

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Post 文章表
type Post struct {
	// gorm.Model
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `gorm:"type:varchar(128);not null" json:"title"`
	Author    string    `gorm:"type:varchar(32);not null" json:"author"`
	Content   string    `gorm:"type:varchar(256)" json:"content"`
	Status    bool      `gorm:"type:boolean;not null" json:"status"`
}

var ormDB = &gorm.DB{}

func init() {
	// go get -u gorm.io/driver/mysql
	// go get -u gorm.io/gorm
	// 连接mysql数据库
	dsn := "luzp:luzp@tcp(127.0.0.1:3306)/vue_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), // 创建日志输出
	})
	if err != nil {
		panic("Failed to connect database")
	}

	// 自动创建表，执行表的迁移操作
	db.AutoMigrate(&Post{})

	ormDB = db
}

// DB 返回连接成功的db
func DB() *gorm.DB {
	return ormDB
}
