package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//对应userinfo数据表
type UserInfo struct {
	ID     uint   //ID默认为主键  其他的需要用primary key
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:4524@(127.0.0.1:3306)/db01?parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("open failed err = ", err)
		return
	}
	defer db.Close()

	//创建表，自动迁移（把结构体和数据表进行对应） 表名默认为结构体名称的复数 
	//禁用表名复数 db.SingularTable(True)
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	// u1 := UserInfo{1, "李白", "男", "写诗"}
	// db.Create(u1)

	//查询

	var u UserInfo
	db.First(&u) //查询表的第一条数据保存在结构体实例u中
	fmt.Printf("查询出来的数据:\n%+v", u)

	//更新
	db.Model(&u).Update("hobby", "喝酒")

	//删除
	db.Delete(&u)

}
