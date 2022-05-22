package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type NewUser struct {
	gorm.Model
	Name string
	Age  int64
}

func showSlcie(a *[]NewUser) {
	for i, v := range *a {
		fmt.Printf("i =%v,v=%+v\n", i, v)
	}
}

func main() {
	//1 连接数据库
	db, err := gorm.Open("mysql", "root:4524@(127.0.0.1:3306)/db01?&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("OPEN MYSQL FAILED ERR=", err)
		return
	}
	defer db.Close()
	//2 自动建表
	db.AutoMigrate(&NewUser{})

	//删除主键字段对应的数据  无主键则会删除全部记录
	// var user NewUser
	// user.ID = 1
	// db.Delete(&user)

	// db.Where("name=?", "大白").Delete(NewUser{})
	// var u []NewUser

	// db.Find(&u)

	// showSlcie(&u)

	//物理删除
	db.Debug().Unscoped().Where("name=?", "小黑").Delete(NewUser{})

}
