package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DatabaseURL    string
	StaticFilePath string

	AdminName     string
	AdminPassword string
)

func init() {
	// 加载配置文件
	err := godotenv.Load("./config/config.env")
	if err != nil {
		log.Fatal("Error loading config.env file", err)
	}

	// 初始化参数
	DatabaseURL = os.Getenv("DATABASE_URL")
	StaticFilePath = os.Getenv("STATIC_FILE_PATH")

	AdminName = os.Getenv("ADMIN_NAME")
	AdminPassword = os.Getenv("ADMIN_PSW")
}
