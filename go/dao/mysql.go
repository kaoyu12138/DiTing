package dao

import "github.com/jinzhu/gorm"

var SqlSession *gorm.DB

func InitMysql() (err error) {
	// 连接MySql数据库
	SqlSession, err = gorm.Open("mysql", "root:12138@tcp(127.0.0.1:3306)/userinfo?charset=utf8mb4&parseTime=True&loc=Local")
	return err
}

func Close() {
	SqlSession.Close()
}
