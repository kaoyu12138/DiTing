package controller

import (
	"DiTing/go/entity"
	"DiTing/go/service"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 输入网址返回分类标签
func GetUrlTag(ctx *gin.Context) {
	Url := ctx.PostForm("url")
	// 判断输入的是否是网址
	_, err := url.ParseRequestURI(Url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断当前网址是否有被数据库记录
	_, ok := service.IsUrlExits(Url)
	if !ok {
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该网址不在数据库中"})
		tag, err3 := service.CmdPythonGetTag(Url)
		if err3 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err3.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取成功", "data": tag})
		}
		return
	}

	tag, err2 := service.GetTag(Url)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取成功", "data": tag})
	}
}

func ShowDangerUrlList(ctx *gin.Context) {
	offsetParam := ctx.Query("offset")
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		offset = 0
	}
	limitParam := ctx.Query("pageSize")
	limit, err2 := strconv.Atoi(limitParam)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post pageSiza"})
		return
	}
	urllist, err3 := service.GetDangerUrlList(offset, limit)
	if err3 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err3.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取成功", "data": urllist})
	}
}

// 用户上传危险网址，判断后存进数据库（或以Tag="unknow"先存进数据库等待判断）
func PostWebsite(ctx *gin.Context) {
	Url := ctx.PostForm("url")
	recordDate := time.Now()

	// 判断输入的是否是网址
	_, err := url.ParseRequestURI(Url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断当前网址是否有被数据库记录
	_, ok := service.IsUrlExits(Url)
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "网址已保存"})
		return
	}
	tag, err3 := service.CmdPythonGetTag(Url)
	if err3 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err3.Error()})
	}
	newWebsite := entity.Website{
		Url:        Url,
		Tag:        tag,
		RecordDate: recordDate,
	}

	err2 := service.CreateWebsite(&newWebsite)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "上传成功"})
	}
}
