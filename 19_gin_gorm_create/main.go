package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type NewUser struct {
	ID   int64
	Name string
	Age  int64
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:4524@(127.0.0.1:3306)/db01?&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("OPEN MYSQL FAILED ERR=", err)
		return
	}
	defer db.Close()

	//结构体和数据库表对应
	db.AutoMigrate(&NewUser{})
	//创建
	u := NewUser{
		Name: "小王子",
		Age:  18,
	}
	isEmpty := db.NewRecord(&u)
	fmt.Println(isEmpty)
	db.Create(&u)
	fmt.Println(db.NewRecord(&u))

}
