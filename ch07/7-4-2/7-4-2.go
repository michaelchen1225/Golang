package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"` //故意將欄位型別改錯
}

func main() {
	p, err := loadPerson("data.json") //讀取同目錄下的文字檔
	if err != nil {
		//若有錯誤, 印出其值和型別
		fmt.Printf("%v", err)
		fmt.Printf("%T", err)
	}
	fmt.Println(p)
}

func loadPerson(fname string) (Person, error) {
	var p Person
	f, err := os.Open(fname)
	if err != nil {
		return p, err //傳回檔案開啟錯誤
	}
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		return p, err //傳回JSON解析錯誤
	}
	return p, err
}
