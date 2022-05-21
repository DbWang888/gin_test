package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/:name/:age", func(c *gin.Context) {
		//获取路径参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month",func(c *gin.Context){
		month := c.Param("month")
		year := c.Param("year")
		c.JSON(http.StatusOK,gin.H{
			"year":year,
			"month":month,
		})
	})

	r.Run(":9090")
}
