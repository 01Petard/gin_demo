package middleware

import (
	"gin_demo/common"
	"gin_demo/model"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authmiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || len(tokenString) < 7 || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(401, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 截取token有效部分
		tokenString = tokenString[7:]

		// 解析token
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{"code": 401, "msg": "token解析失败，权限不足"})
			ctx.Abort()
			return
		}

		// 获取用户id
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 如果用户不存在
		if user.ID == 0 {
			ctx.JSON(401, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//如果用户存在满了user信息写入上下文
		ctx.Set("user", user)

		ctx.Next()

	}
}
