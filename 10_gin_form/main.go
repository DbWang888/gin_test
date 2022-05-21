package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html","./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})


	r.POST("/login", func (c *gin.Context)  {
		//获取form表单参数，其余两种为c.GetPostForm和c.DefaultPostForm
		//用法和query一致
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(http.StatusOK,"index.html",gin.H{
			"Name":username,
			"Password":password,
		})
	})


	r.Run(":9090")
}
