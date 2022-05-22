package models

import "Golang_gin_test/24_gin_exec_re/dao"

//model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	todo这个model增删改查都放在这里
*/

//向数据库添加数据
func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

//查询获取所有数据
func GetTodos() (todoList []*Todo, err error) {
	if err = dao.DB.Find(todoList).Error; err != nil {
		return nil, err
	}
	return
}

//修改数据
func ModifyTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return

}

//删除数据
func DelTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
