package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//使用gin.H拼接json格式字符串
func getJson1(c *gin.Context) {
	//方式一 ： 自己拼接json格式字符串
	c.JSON(http.StatusOK, gin.H{"name": "jack"})
}

// 直接使用结构体返回json格式字符串
func getJson2(c *gin.Context) {
	var mes struct {
		Name   string
		Age    int
		Gender string
	}

	mes.Name = "小王子"
	mes.Age = 19
	mes.Gender = "男"
	c.JSON(http.StatusOK, mes)
}

func main() {
	r := gin.Default()

	r.GET("/someJson", getJson1)
	r.GET("/some", getJson2)

	r.Run(":8080")
}
