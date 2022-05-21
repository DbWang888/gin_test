package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	//主模板在前
	if err != nil {
		fmt.Println("ParseFiles failed err = ", err)
		return
	}
	name := "大王"

	tmpl.Execute(w, name)

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("listenAndService is failed err = ", err)
		return
	}

}
