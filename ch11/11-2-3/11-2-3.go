package main

import (
	"encoding/json"
	"fmt"
)

type person struct { //父結構
	Lastname  string  `json:"lname"`
	Firstname string  `json:"fname"`
	Address   address `json:"address"` //子結構欄位型別
}

type address struct { //子結構
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

func main() {
	//JSON資料
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
	//解析JSON並將值存入結構
	p := person{}
	if err := json.Unmarshal(data, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", p)
}
