package main

import (
	"log"
	"net/http"
)

func main() {
	//直接將一個匿名函式傳給HandleFunc
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//把index.html當成回應寫入ResponseWriter
		http.ServeFile(w, r, "./index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
