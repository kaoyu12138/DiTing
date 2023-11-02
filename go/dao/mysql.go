package dao

import "github.com/jinzhu/gorm"

var SqlSession *gorm.DB

func InitMysql() (err error) {
	// 连接MySql数据库   // !!!!!!!!!!!!!! 密码改成自己本地数据库密码
	SqlSession, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/userinfo?charset=utf8mb4&parseTime=True&loc=Local")
	return err
}

func Close() {
	SqlSession.Close()
}
