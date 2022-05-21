package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	//1 编写模板
	//2 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parsrfiles err = ", err)
		return
	}

	//3 渲染模板
	user := User{
		Name:   "小明",
		Age:    18,
		Gender: "男",
	}

	//此处可传入结构体、map、普通变量
	err = t.Execute(w, user)

	if err != nil {
		fmt.Println("excute err = ", err)
		return
	}

}

func rangeTest(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./range.tmpl")
	if err != nil {
		fmt.Println("parsrfiles err = ", err)
		return
	}

	//3 渲染模板
	array := [3]int{1, 2, 3}

	//此处可传入结构体、map、普通变量
	err = t.Execute(w, array)

	if err != nil {
		fmt.Println("excute err = ", err)
		return
	}
}

func main() {
	http.HandleFunc("/", rangeTest)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed err = ", err)
		return
	}

}
