package main

import (
	"log"
	"net/http"
)

type hello struct{} //HTTP請求處理器

//請求處理器的方法實作
func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>" //有HTML標籤的文字
	w.Write([]byte(msg))          //寫入回應 (傳給客戶端)
}

func main() {
	//啟動服務
	log.Fatal(http.ListenAndServe(":8080", hello{}))
}
