package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// r.MaxMultipartMemory 最大内存限制，超出内存，自动保存在系统临时文件中
	//再保存
	r.POST("/upload", func(c *gin.Context) {
		//处理上传请求，
		//从请求中读取文件，并保存到本地（服务端本地）
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			// dst := fmt.Sprintln("./%s",f.Filename)
			dst := path.Join("./", f.Filename)
			err := c.SaveUploadedFile(f, dst)
			if err != nil {
				fmt.Println("savefile err = ", err)
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.Run(":8080")
}
