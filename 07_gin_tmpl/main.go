package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//加载静态文件，入css,js等
	r.Static("/xxx", "./statics") //指向以XXX开头的文件夹的static文
	//gin框架中添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	// r.LoadHTMLFiles("./index.tmpl") //模板解析
	//匹配正则表达式通配符
	r.LoadHTMLGlob("templates/**/*")

	//访问/index时执行后面匿名函数
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ //模板渲染
			"title": "gin模板posts",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ //模板渲染
			"title": "gin模板users,<a href='http://www.baidu.com'>百度</a>",
		})
	})

	//下载引入的模板
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/home.html", nil)
	})

	r.Run(":8080") //启动Server
}
