package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//定义模型
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	Adress       string  `gorm:"index;addr"`
	IgnoreMe     int     `gorm:"-"`
}

func main() {
	db, err := gorm.Open("mysql", "root:4524@(127.0.0.1:3306)/db01?&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("OPEN MYSQL FAILED ERR=", err)
		return
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	//使用user结构体创建一个表名为的Mall_User表
	db.Table("Mall_User").CreateTable(&User{})

}
