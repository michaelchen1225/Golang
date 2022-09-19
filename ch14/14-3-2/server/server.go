package main

import (
	"log"
	"net/http"
)

type server struct{} //伺服器結構

//收到請求時要執行的伺服器服務
func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := `{"message": "hello world"}` //要傳回給客戶端的JSON資料
	w.Write([]byte(msg))                //將JSON字串寫入回應主體
}

func main() {
	//啟動本地端伺服器, 監聽port 8080(即http://localhost:8080)
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
