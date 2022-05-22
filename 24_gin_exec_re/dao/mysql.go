package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//定义全局变量gorm
var (
	DB *gorm.DB
)

//连接MySql数据库并测试联通性
func InitMysql() (err error) {
	dsn := "root:4524@(127.0.0.1:3306)/bubble?&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = DB.DB().Ping()
	return
}
