package models

import (
	"embedded-album/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"time"
)

var DB *gorm.DB

func init() {
	// 连接数据库
	db, err := gorm.Open("sqlite3", config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	// 连接参数
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	// Migrate the schema
	db.AutoMigrate(&Picture{})

	// 全局赋值
	DB = db
}
