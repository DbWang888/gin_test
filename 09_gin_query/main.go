package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//获取浏览器发请求携带的query string参数
		//key=value格式，多个key-value在浏览器内用&连接
		name := c.Query("query")
		age := c.Query("age")
		// name := c.DefaultQuery("query", "somebody")  //根据query取值，取不到则默认为somebody
		// name, ok := c.GetQuery("query")
		// if !ok {
		// 	name = "取不到"
		// }
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})

	})
	r.Run(":8080")
}
