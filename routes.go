package main

import (
	"gin_demo/controller"
	"gin_demo/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) *gin.Engine {
	// 注册
	r.POST("/api/auth/register", controller.Register)
	// 登录
	r.POST("/api/auth/login", controller.Login)
	// 获取用户信息
	r.GET("/api/auth/info", middleware.Authmiddleware(), controller.Info)
	return r
}
