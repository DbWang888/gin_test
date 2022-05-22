package routers

import (
	"Golang_gin_test/24_gin_exec_re/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.InderHandler)

	//v1路由组
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreatedTodo)
		//查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 查看所有待办事项
		v1Group.GET("/todo", controller.GetAllTodo)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r

}
