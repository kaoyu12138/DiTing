package controller

import (
	"DiTing/go/entity"
	"DiTing/go/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
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
		ctx.JSON(200, gin.H{"code": 200, "msg": "登录成功"})
	}
}
