package controller

import (
	"Golang_gin_test/24_gin_exec_re/dao"
	"Golang_gin_test/24_gin_exec_re/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InderHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreatedTodo(c *gin.Context) {
	//前端页面发请求至此
	//把前端页面请求的数据放入数据库
	//返回响应
	var todo models.Todo
	c.BindJSON(&todo)
	err := dao.DB.Create(&todo).Error
	if err != nil {
		fmt.Println("DB.create failed err = ", err)
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

//获取全部
func GetAllTodo(c *gin.Context) {
	var todos []models.Todo
	err := dao.DB.Find(&todos).Error
	if err != nil {
		fmt.Println("DB.Find failed err=", err)
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todos)
	}

}

//修改todo
func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		fmt.Println("get did failed")
		return
	}

	var todo models.Todo
	err := dao.DB.Where("id=?", id).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
	c.BindJSON(&todo)
	err = dao.DB.Save(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//删除todo
func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		fmt.Println("get did failed err=")
		return
	}
	dao.DB.Where("id=?", id).Delete(&models.Todo{})
}
