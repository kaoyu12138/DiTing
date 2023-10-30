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
		DiTingGroup.POST("/user/register", controller.Register)
		DiTingGroup.POST("/user/login", controller.Login)
		DiTingGroup.POST("/user/avatar", controller.PostUserAvatar)
		DiTingGroup.GET("/user/message", controller.GetUserMessage)
		DiTingGroup.GET("/user/showEssay", controller.GetUserEssay)

		DiTingGroup.POST("/essay/post", controller.PostEssay)
		DiTingGroup.GET("/essay/show", controller.ShowEssay)
		DiTingGroup.POST("/essay/likes:id", controller.UpdateLikeCount)
	}

	return r
}
