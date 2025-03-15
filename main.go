package main

import (
	"gin_demo/common"
	"github.com/gin-gonic/gin"
)

func main() {

	common.InitDB()

	r := gin.Default()
	r = SetupRoutes(r)

	panic(r.Run(":8080"))
}
