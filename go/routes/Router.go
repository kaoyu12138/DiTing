package routes

import (
	"DiTing/go/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	//创建各api的路由
	DiTingGroup := r.Group("DiTing")
	{
		DiTingGroup.POST("/register", controller.Register)
		DiTingGroup.POST("/login", controller.Login)
	}

	return r
}
