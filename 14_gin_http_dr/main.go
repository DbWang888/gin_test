package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//重定向

func main() {

	r := gin.Default()

	// r.GET("/test", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": "ok",
	// 	})
	// })

	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com") //跳转别的连接
	})

	r.GET("/a", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"url":     "a",
		})
	})

	r.GET("/b", func(c *gin.Context) {
		c.Request.URL.Path = "/a" //把请求的url修改,类似转发
		r.HandleContext(c)
	})

	r.Run(":8080")
}
