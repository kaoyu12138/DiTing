package main

import (
	"DiTing/go/dao"
	"DiTing/go/entity"
	"DiTing/go/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	defer dao.Close()
	//绑定模型
	dao.SqlSession.AutoMigrate(&entity.User{})
	dao.SqlSession.AutoMigrate(&entity.Essay{})
	dao.SqlSession.AutoMigrate(&entity.Website{})
	//注册路由
	r := routes.SetRouter()
	//启动端口为8081的项目
	r.Run(":8081")
}
