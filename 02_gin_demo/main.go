package main

import (
	"github.com/gin-gonic/gin"
)

func sayhello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "func say hello",
	})

}
func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	//GET请求方式；/hello:请求的路径
	//当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	// r.GET("/hello", func(ctx *gin.Context) {
	// 	//ctx.JSON:返回json格式的数据
	// 	ctx.JSON(200, gin.H{
	// 		"message": "hello,word!",
	// 	})
	// })
	r.GET("/hello", sayhello)
	//启动http服务，默认在0.0.0.0：8080启动服务
	r.Run()
}
