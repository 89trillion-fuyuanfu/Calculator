package route

import (
	"Calculator/ctrl"
	"github.com/gin-gonic/gin"
)

func Getroute() {

	// 路由配置
	router := gin.Default()
	port := "8080"
	//匹配calculate?name=xxx
	router.POST("/calculate", ctrl.Getstring)

	router.Run(":" + port)
}
