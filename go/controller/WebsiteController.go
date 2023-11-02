package controller

import (
	"DiTing/go/entity"
	"DiTing/go/service"
	"net/http"
	"net/url"
	"strconv"

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
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该网址不在数据库中"})
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
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 0
	}
	urllist, err2 := service.GetDangerUrlList(offset, limit)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取成功", "data": urllist})
	}
}

// 用户上传危险网址，等待判断后存进数据库（或以Tag="unknow"先存进数据库等待判断）
func PostWebsite(ctx *gin.Context) {
	Url := ctx.PostForm("url")
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

	newWebsite := entity.Website{
		Url: Url,
		Tag: "unknow",
	}

	err = service.CreateWebsite(&newWebsite)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "上传成功"})
	}
}
