package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:"message"`
}

func postDataAndReturnResponse(msg messageData) messageData {
	jsonBytes, _ := json.Marshal(msg)    //將要傳的結構編碼成JSON
	buffer := bytes.NewBuffer(jsonBytes) //將JSON轉成bytes.Buffer

	//送出POST請求和資料 , 標頭為application/json, 並接收回應
	r, err := http.Post("http://localhost:8080", "application/json", buffer)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	message := messageData{}
	//解碼伺服器回應的JSON資料
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Fatal(err)
	}

	return message
}

func main() {
	msg := messageData{Message: "hi server!"} //要傳給伺服器的訊息
	data := postDataAndReturnResponse(msg)    //接收伺服器串回的訊息

	fmt.Println(data.Message)
}
