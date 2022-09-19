package main

import (
	"encoding/json"
	"fmt"
)

type book struct {
	ISBN          string `json:"-"` //短折線, 代表直接略過
	Title         string `json:"title"`
	YearPublished int    `json:"yearpub,omitempty"`
	Author        string `json:""`           //沒有鍵名稱, 代表加入欄位並沿用欄位名稱
	CoAuthor      string `json:",omitempty"` //沒有鍵名稱
}

func main() {
	b := book{}
	b.ISBN = "9933HIST" //由於已指定略過, 不會出現在JSON中
	b.Title = "Greatest of all Books"
	b.YearPublished = 2020
	b.Author = "John Adams"
	//沒對CoAuthur賦值, 所以會因為omitempty的原因被略過

	json, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}
