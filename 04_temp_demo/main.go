package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//模板自定义函数

func sayHello(w http.ResponseWriter, r *http.Request) {
	// stringName := "模板同志"

	kua := func(arg string) (string, error) {
		return arg + "ok", nil
	}

	//采用链式操作在Parse之前调用Funcs添加自定义函数
	tmpl, err := template.New("temp.tmpl").Funcs(template.FuncMap{"kua": kua}).ParseFiles("./temp.tmpl")
	//此处New（）函数内的名称需要与ParseFiles内的一致
	if err != nil {
		fmt.Println("Funcs failed err = ", err)
		return
	}
	tmpl.Execute(w, "李白")

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("listenAndService is failed err = ", err)
		return
	}

}
