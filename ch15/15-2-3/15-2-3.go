package main

import (
	"log"
	"net/http"
)

type hello struct{}

//原本的請求處理器方法
func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>"
	w.Write([]byte(msg))
}

//新函式, 用來處理對路徑 /page1的請求
func servePage1(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Page 1</h1>"
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/page1", servePage1)
	http.Handle("/", hello{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
