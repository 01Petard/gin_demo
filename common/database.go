package common

import (
	"fmt"
	"gin_demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	host := "1.94.147.176"
	port := "3306"
	database := "gin_demo"
	username := "root"
	password := "kjiolluy711"
	charset := "utf8mb4"

	// MySQL 连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, database, charset)

	// 连接 MySQL
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 自动迁移
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	log.Println("数据库连接成功")
}

func GetDB() *gorm.DB {
	return DB
}
