package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"page": "home",
		})
	})

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"page": "index",
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"page": "login",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.Request.URL.Path = ("/index")
		r.HandleContext(c)
	})

	r.Run(":8080")
}
