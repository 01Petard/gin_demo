package controller

import (
	"gin_demo/common"
	"gin_demo/model"
	"gin_demo/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()

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
	if isPhoneExist(DB, phone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}

	// 创建新用户
	// 如果没有名称，则生成一个10为随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}

	DB.Create(&newUser)

	log.Println(name, phone, password)

	ctx.JSON(http.StatusOK, gin.H{
		"mag": "注册成功！",
	})
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
