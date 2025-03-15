package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func main() {

	db := InitDB()

	r := gin.Default()

	r.POST("/api/auth/register", func(ctx *gin.Context) {

		name := ctx.PostForm("name")
		phone := ctx.PostForm("phone")
		password := ctx.PostForm("password")

		if len(phone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}

		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}

		// 判断手机号是否存在，用户是否存在
		if isPhoneExist(db, phone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}

		// 创建新用户
		// 如果没有名称，则生成一个10为随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}
		newUser := User{
			Name:     name,
			Phone:    phone,
			Password: password,
		}

		db.Create(&newUser)

		log.Println(name, phone, password)

		ctx.JSON(http.StatusOK, gin.H{
			"mag": "注册成功！",
		})

	})

	panic(r.Run(":8080"))
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
func InitDB() *gorm.DB {
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	db.AutoMigrate(&User{})

	return db
}
