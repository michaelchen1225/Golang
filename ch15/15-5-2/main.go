package main

import (
	"log"
	"net/http"
)

func main() {

	//對任何路徑提供index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	//將/statics路徑對應到本地的/public資料夾
	http.Handle(
		"/statics/",
		http.StripPrefix("/statics/", http.FileServer(http.Dir("./public"))),
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
