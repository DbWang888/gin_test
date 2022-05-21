package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		// username := c.Query("username")
		// password := c.Query("pwd")

		// u := UserInfo{
		// 	username: username,
		// 	password: password,
		// }

		var u UserInfo
		err := c.ShouldBind(&u)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "err",
			})
		} else {
			fmt.Println("username = ", u.Username)
			fmt.Println("password = ", u.Password)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}

	})

	r.Run(":9090")
}
