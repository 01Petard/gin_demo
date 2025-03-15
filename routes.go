package main

import (
	"gin_demo/controller"
	"gin_demo/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.Authmiddleware(), controller.Info)
	return r
}
