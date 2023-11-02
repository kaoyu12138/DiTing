package controller

import (
	"DiTing/go/entity"
	"DiTing/go/service"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	//获取前端传来的注册人信息
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//判断手机号码不正确，或密码少于六位，返回错误消息给前端，并返回
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	//判断当前手机号码是否已经存在
	_, ok := service.IsTelephoneExits(telephone)
	if ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "电话已经存在"})
		return
	}
	_, ok = service.IsUserNameExits(name)
	if ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户名已经存在"})
		return
	}

	//对密码进行加密，加密后存入数据库user表
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	newUser := entity.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}

	err := service.CreateUser(&newUser)
	//判断是否异常，无异常则返回包含200和注册成功的信息
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
	}
}

func Login(ctx *gin.Context) {
	//使用telephone + password登录，从前端获取
	userName := ctx.PostForm("userName")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//判断密码少于六位，返回错误消息给前端，并返回
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	//判断当前手机号是否在数据库中存在，若存在则取出该user体
	user, ok := service.IsTelephoneExits(telephone)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	//对比输入的password和数据库中该user的password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	} else {
		//创建一个session会话容器，并以键值对的方式保存当前会话的telephone
		s := sessions.Default(ctx)
		s.Set("telephone", telephone)
		s.Set("userName", userName)
		s.Save()

		ctx.JSON(200, gin.H{"code": 200, "msg": "登录成功"})
	}
}

func PostUserAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile(("avatar"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, reader, fileSize, &putExtra)
	if err != nil {
		fmt.Println("formUploader.PutWithoutKey err: ", err)
	}
	url := ImgUrl + ret.Key

	telephone := sessions.Default(ctx).Get("telephone")
	if err = service.UpdateAvatar(telephone.(string), url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "上传失败"})
	} else {
		s := sessions.Default(ctx)
		s.Set("userAvatar", url)
		s.Save()

		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "上传头像成功", "data": url})
	}
}

func GetUserMessage(ctx *gin.Context) {
	telephone := sessions.Default(ctx).Get("telephone")
	var user entity.User
	err := service.GetUser(telephone.(string), &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取成功", "data": user})
	}
}

func GetUserEssay(ctx *gin.Context) {
	userName := sessions.Default(ctx).Get("userName")
	offsetParam := ctx.Query("offset")
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		offset = 0
	}

	essayList, err2 := service.GetUserEssayList(offset, userName.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取成功", "data": essayList})
	}
}
