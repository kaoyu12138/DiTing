package routes

import (
	"DiTing/go/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))

	//创建各api的路由
	DiTingGroup := r.Group("DiTing")
	{
		DiTingGroup.POST("/user/register", controller.Register)
		DiTingGroup.POST("/user/login", controller.Login)
		DiTingGroup.POST("/user/avatar", controller.PostUserAvatar)
		DiTingGroup.GET("/user/message", controller.GetUserMessage)
		DiTingGroup.GET("/user/showEssay", controller.GetUserEssay)

		DiTingGroup.POST("/essay/post", controller.PostEssay)
		DiTingGroup.GET("/essay/showList", controller.ShowEssayList)
		DiTingGroup.GET("/essay/showEssay", controller.ShowEssay)
		DiTingGroup.GET("/essay/likes", controller.UpdateLikeCount)

		DiTingGroup.POST("/website/post", controller.PostWebsite)
		DiTingGroup.GET("/website/tag", controller.GetUrlTag)
		DiTingGroup.GET("/website/showList", controller.ShowDangerUrlList)

	}

	return r
}
