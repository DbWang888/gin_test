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

	//更新所有字段
	var user NewUser
	db.First(&user)
	// user.Age = 60
	// user.Name = "老白"
	// db.Debug().Save(&user)

	//更新指定字段
	// db.Debug().Model(&user).Update("name", "张三")

	// 更新选定字段  select指定某个字段  omit忽略某个字段

	// m1 := map[string]interface{}{
	// 	"name": "马六",
	// 	"age":  43,
	// }
	// db.Model(&user).Updates(m1)  //m1列出来的所有字段都会更新
	// db.Model(&user).Select("name").Update(m1) //只更新name字段
	// db.Model(&user).Omit("name").Update(m1)  //忽略更新name字段

	//不更新钩子函数对应的时间等字段
	// db.Debug().Model(&user).UpdateColumns(m1)

	//批量处理，所有数据的age字段增加5

	db.Model(&NewUser{}).Update("age", gorm.Expr("age + ?", "5"))

}
