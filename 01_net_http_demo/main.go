package main

//使用net/http标准库实现web开发
import (
	"fmt"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "<h1>hello,golang!</h1>")
}

func main() {
	http.HandleFunc("/hello", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed err = ", err)
	}
}
