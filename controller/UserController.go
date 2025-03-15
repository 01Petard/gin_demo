package controller

import (
	"gin_demo/common"
	"gin_demo/dto"
	"gin_demo/model"
	"gin_demo/response"
	"gin_demo/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()

	// 获取参数
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(phone) != 11 {
		response.Response(
			ctx,
			http.StatusUnprocessableEntity,
			422,
			gin.H{}, "手机号必须为11位",
		)
		return
	}
	if len(password) < 6 {
		response.Response(
			ctx,
			http.StatusUnprocessableEntity,
			422,
			gin.H{}, "密码不能少于6位",
		)
		return
	}

	// 判断手机号是否存在，用户是否存在
	var user model.User
	DB.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		response.Response(
			ctx,
			http.StatusUnprocessableEntity,
			422,
			gin.H{}, "用户已经存在",
		)
		return
	}

	// 创建新用户
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(
			ctx,
			http.StatusInternalServerError,
			500,
			gin.H{}, "加密错误！",
		)
	}

	// 如果没有名称，则生成一个10为随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hashedPassword),
	}

	DB.Create(&newUser)

	log.Println(name, phone, password)

	response.Success(ctx, gin.H{
		"code": "200",
		"data": gin.H{},
		"mag":  "登录成功！",
	}, "注册成功！")
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()

	// 获取参数
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(phone) != 11 {
		response.Response(
			ctx,
			http.StatusUnprocessableEntity,
			422,
			gin.H{}, "手机号必须为11位",
		)
		return
	}
	if len(password) < 6 {
		response.Response(
			ctx,
			http.StatusUnprocessableEntity,
			422,
			gin.H{}, "密码不能少于6位",
		)
		return
	}

	// 判断手机号是否存在，用户是否存在
	var user model.User
	DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		response.Response(
			ctx,
			http.StatusUnprocessableEntity,
			422,
			gin.H{}, "用户不存在，请先注册！",
		)
		return
	}

	// 判断密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.Response(
			ctx,
			http.StatusBadRequest,
			400,
			gin.H{}, "密码错误！")
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(
			ctx,
			http.StatusInternalServerError,
			500,
			gin.H{}, "系统错误！",
		)
		log.Printf("token generate error: %v", err)
		return
	}

	response.Success(
		ctx,
		gin.H{
			"code": "200",
			"data": gin.H{
				"token": token,
			},
			"mag": "登录成功！",
		},
		"登录成功！")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(
		ctx,
		gin.H{
			"code": "200",
			"data": gin.H{
				"user": dto.ToUserDto(user.(model.User))},
		},
		"获取用户信息成功！",
	)

}
