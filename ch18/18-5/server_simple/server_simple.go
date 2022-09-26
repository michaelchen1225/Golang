package main

import (
	"log"
	"net/http"
)

const ( //憑證及私鑰檔名
	serverCertName = `.\server_cert.pem`
	serverKeyName  = `.\server.key`
)

func main() {
	log.Println("啟動伺服器")
	//對路徑指定請求處理函式
	http.HandleFunc("/", hello)
	//啟動HTTPS/TLS伺服器, 仔入憑證和私鑰
	log.Fatal(http.ListenAndServeTLS(":8080", serverCertName, serverKeyName, nil))
}

//HTTP請求處理函式
func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("收到請求")
	w.Write([]byte("Hello Golang from a secure server"))
}
