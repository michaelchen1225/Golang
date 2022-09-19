package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getDataAndReturnResponse() string {
	//送出GET請求和取得會應
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() //在結束時釋放r占用的連線資源

	if resp.StatusCode != http.StatusOK { //檢查HTTP狀態碼
		log.Fatal(resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data) //回傳回應內容
}

func main() {
	data := getDataAndReturnResponse()
	fmt.Print(data) //印出回應內容
}
