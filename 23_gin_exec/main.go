package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//定义全局变量gorm
var (
	DB *gorm.DB
)

//model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//连接MySql数据库并测试联通性
func initMysql() (err error) {
	dsn := "root:4524@(127.0.0.1:3306)/bubble?&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = DB.DB().Ping()
	return
}

func main() {

	//1 联通数据库
	err := initMysql()
	if err != nil {
		fmt.Println("mysql failed err = ", err)
		return
	}
	defer DB.Close()

	//2 模型绑定
	DB.AutoMigrate(&Todo{}) //创建的表名为todos

	//3 gin框架初始化
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//v1路由组
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面发请求至此
			//把前端页面请求的数据放入数据库
			//返回响应
			var todo Todo
			c.BindJSON(&todo)
			err = DB.Create(&todo).Error
			if err != nil {
				fmt.Println("DB.create failed err = ", err)
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})
		//查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todos []Todo
			err = DB.Find(&todos).Error
			if err != nil {
				fmt.Println("DB.Find failed err=", err)
				c.JSON(http.StatusOK, gin.H{
					"err": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todos)
			}

		})
		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				fmt.Println("get did failed err=", err)
				return
			}

			var todo Todo
			err = DB.Where("id=?", id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err": err.Error(),
				})
			}
			c.BindJSON(&todo)
			err = DB.Save(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				fmt.Println("get did failed err=", err)
				return
			}
			DB.Where("id=?", id).Delete(&Todo{})
		})
	}

	r.Run(":9999")
}
