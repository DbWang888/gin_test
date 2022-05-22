package main

import (
	"Golang_gin_test/24_gin_exec_re/dao"
	"Golang_gin_test/24_gin_exec_re/models"
	"Golang_gin_test/24_gin_exec_re/routers"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	//1 联通数据库
	err := dao.InitMysql()
	if err != nil {
		fmt.Println("mysql failed err = ", err)
		return
	}
	defer dao.DB.Close()

	//2 模型绑定
	dao.DB.AutoMigrate(&models.Todo{}) //创建的表名为todos

	//3 gin框架初始化
	r := routers.SetupRouter()

	r.Run(":9999")
}
