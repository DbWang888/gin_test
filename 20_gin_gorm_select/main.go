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

func bai(db *gorm.DB) *gorm.DB {
	return db.Where("name LIKE ?", "%白%")
}

func ageMoreThan30(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", "30")
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

	// u1 := NewUser{
	// 	Name: "宋江",
	// 	Age:  34,
	// }
	// //3 插入
	// db.Create(&u1)

	//4 查询
	//var user NewUser
	//4.1
	//db.First(&user) //根据主键查询，且主键为数字

	//fmt.Printf("第一条内容是：%+v", user)

	//查询所有信息，以切片形式呈现
	//4.2
	//var users []NewUser
	//db.Debug().Find(&users)
	// for i, v := range users {
	// 	fmt.Printf("i =%v,v=%+v\n", i, v)
	// }

	//where条件查询
	// 1
	// var users []NewUser
	// db.Where("name = ?", "李白").Find(&users)
	// for i, v := range users {
	// 	fmt.Printf("i =%v,v=%+v\n", i, v)
	// }

	// 2
	// var users []NewUser
	// db.Where("name <> ?", "李白").Find(&users)
	// showSlcie(&users)

	// 3
	// var users []NewUser
	// db.Where("name LIKE ?", "%白%").Find(&users)
	// showSlcie(&users)

	//4
	var users []NewUser
	db.Scopes(bai, ageMoreThan30).Find(&users)
	showSlcie(&users)

}
