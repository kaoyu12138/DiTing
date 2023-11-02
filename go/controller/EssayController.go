package controller

import (
	"DiTing/go/entity"
	"DiTing/go/service"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func PostEssay(ctx *gin.Context) {
	essayName := ctx.PostForm("essayName")
	essayContent := ctx.PostForm("essayContent")
	publishDate := time.Now()

	file, _ := ctx.FormFile(("essayAvatar"))
	fileSize := file.Size
	f, _ := file.Open()
	buf := make([]byte, file.Size)
	f.Read(buf)
	var AccessKey = "2w2sKxZdcbcTdad4HAetvTtttxyeL_cGmyDMv6WE"
	var SerectKey = "DTdvfx0zj_hooKKcqOZmYgdxbPZvAMREuBUrURGC"
	var Bucket = "diting"                       // 前边创建的空间名称
	var ImgUrl = "s3bgzzv3a.hd-bkt.clouddn.com" // 前边给的测试域名
	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SerectKey)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	reader := bytes.NewReader(buf)
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, reader, fileSize, &putExtra)
	if err != nil {
		fmt.Println("formUploader.PutWithoutKey err: ", err)
	}
	url := ImgUrl + ret.Key

	name := sessions.Default(ctx).Get("userName")
	avatar := sessions.Default(ctx).Get("userAvatar")

	newEssay := entity.Essay{
		LikeCount:    0,
		UserName:     name.(string),
		UserAvatar:   avatar.(string),
		EssayName:    essayName,
		EssayContent: essayContent,
		EssayAvatar:  url,
		PublishDate:  publishDate,
	}

	err = service.CreateEssay(&newEssay)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "发表成功"})
	}
}

func ShowEssayList(ctx *gin.Context) {
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

func ShowEssay(ctx *gin.Context) {
	idParam := ctx.Query("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	essayContent, err2 := service.GetEssay(id)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "查看成功", "data": essayContent})
	}
}

func UpdateLikeCount(ctx *gin.Context) {
	idParam := ctx.Query("id")
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
