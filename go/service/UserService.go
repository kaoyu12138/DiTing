package service

import (
	"DiTing/go/dao"
	"DiTing/go/entity"
)

// IsTelephoneExits 判断一串手机号是否存在于数据库
func IsTelephoneExits(telephone string) (entity.User, bool) {
	var user entity.User
	dao.SqlSession.Where("Telephone = ?", telephone).Find(&user)
	if user.Id != 0 {
		return user, true
	}
	return user, false
}

func IsUserNameExits(userName string) (entity.User, bool) {
	var user entity.User
	dao.SqlSession.Where("UserName = ?", userName).Find(&user)
	if user.Id != 0 {
		return user, true
	}
	return user, false
}

func CreateUser(user *entity.User) (err error) {
	if err = dao.SqlSession.Create(user).Error; err != nil {
		return err
	}
	return
}

func UpdateAvatar(telephone, avatar string) (err error) {
	if err = dao.SqlSession.Model(new(entity.User)).Where("Telephone = ?", telephone).Update("Avatar", avatar).Error; err != nil {
		return err
	}
	return
}

func GetUser(telephone string, user *entity.User) (err error) {
	if err = dao.SqlSession.Where("Telephone = ?", telephone).Find(user).Error; err != nil {
		return err
	}
	return
}
