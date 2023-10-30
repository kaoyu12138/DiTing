package controller

import (
	"DiTing/go/entity"
	"DiTing/go/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func PostEssay(ctx *gin.Context) {
	essayName := ctx.PostForm("essayName")
	essayContent := ctx.PostForm("essayContent")
	telephone := sessions.Default(ctx).Get("telephone")
	publishDate := time.Now()

	var user *entity.User
	_ = service.GetUser(telephone.(string), user)

	newEssay := entity.Essay{
		LikeCount:    0,
		UserName:     user.Name,
		UserAvatar:   user.Avatar,
		EssayName:    essayName,
		EssayContent: essayContent,
		PublishDate:  publishDate,
	}

	err := service.CreateEssay(&newEssay)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "发表成功"})
	}
}

func ShowEssay(ctx *gin.Context) {
	offsetParam := ctx.Query("offset")
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		offset = 0
	}

	essayList, err2 := service.GetEssayList(offset)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取成功", "data": essayList})
	}
}

func UpdateLikeCount(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	likeCount, err2 := service.UpdateLikeCount(id)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "点赞成功", "data": likeCount})
	}
}
