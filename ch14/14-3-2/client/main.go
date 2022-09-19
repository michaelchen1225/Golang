package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type messageData struct { //對應JSON資料的結構
	Message string `json:"message"`
}

func getDataAndReturnResponse() messageData {
	//對http://localhost:8080送出GET請求
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	//從回應主體讀出所有內容字串
	data, err := io.ReadAll(r.Body)
	//解析JSON字串並將資料存入\message(messageData結構)
	if err != nil {
		log.Fatal(err)
	}

	message := messageData{}
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}

	return message
}

func main() {
	data := getDataAndReturnResponse()
	fmt.Println(data.Message) //印出message欄位的字串
}
