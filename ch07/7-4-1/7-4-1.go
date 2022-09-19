package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Person struct { //用於JSON資料的結構
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	s := `{"Name":"Joe","Age":18}`   //第一筆資料
	s2 := `{"Name":"Jane","Age":21}` //第二筆資料

	//第一筆資料 (字串)
	p, err := loadPerson(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)

	//第二筆資料
	//strings.NewReader()會回傳一個strings.Reader結構,符合io.Reader介面
	p2, err := loadPerson2(strings.NewReader(s2))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)

	//第三筆資料 (讀取檔案後傳回os.File結構,符合io.Reader介面)
	f, err := os.Open("data.json") //開啟同資料夾下的文字檔
	if err != nil {
		fmt.Println(err)
	}
	p3, err := loadPerson2(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p3)
}

//第一個JSON解析函式,接收字串參數
func loadPerson(s string) (Person, error) {
	var p Person
	err := json.NewDecoder(strings.NewReader(s)).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}

//第二個JSON解析函式,接收io.Reader介面參數
func loadPerson2(r io.Reader) (Person, error) {
	var p Person
	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, err
}
