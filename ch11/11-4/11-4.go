package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	Lastname  string  `json:"lname"`
	Firstname string  `json:"fname"`
	Address   address `json:"address"`
}

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

func main() {
	data := []byte(`
{
	"lname":"Smith",
	"fname":"John",
	"address":{
		"street":"Sulphur Springs Rd",
		"city":"Park City",
		"state":"VA",
		"zipcode":12345
	}
}	 
`)
	dataStr := string(data)
	p := person{}

	//用strings.NewReader()從字串建立一個io.Reader
	//並以此建立json.Decoder
	decoder := json.NewDecoder(strings.NewReader(dataStr))
	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(p)
	fmt.Println()

	//建立json.Encoder, 寫入對象是os.Stdout (主控台)
	encoder := json.NewEncoder(os.Stdout)
	//設定前綴詞和縮排文字
	encoder.SetIndent("", "\t")
	//將結構p編碼成JSON
	if err := encoder.Encode(p); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
