package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/base.tmpl", "./template/index.tmpl")

	if err != nil {
		fmt.Println("index parse failed err = ", err)
		return
	}
	msg := "这是index页面"
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/base.tmpl", "./template/home.tmpl")

	if err != nil {
		fmt.Println("home parse failed err = ", err)
		return
	}
	msg := "这是home页面"
	t.Execute(w, msg)

}

func main() {

	http.HandleFunc("/home", home)
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("listen is failed err = ", err)
		return
	}

}
