package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"page": "index",
	})
	fmt.Println("index")
}

//中间件
func test(c *gin.Context) {
	fmt.Println("test in-----")

	start := time.Now()
	c.Next() //调用后续的处理函数
	useTime := time.Since(start).Nanoseconds()

	fmt.Println("耗时", useTime)
	fmt.Println("test out-----")

}

func main() {
	r := gin.Default()

	r.Use(test) //全局注册中间件函数

	// r.GET("/index", test, index)
	r.GET("/index", index)

	r.Run()
}
