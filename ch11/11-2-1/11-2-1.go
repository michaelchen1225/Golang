package main

import (
	"encoding/json"
	"fmt"
)

type greeting struct { //用來儲存JSON資料的結構
	Message string
}

func main() {
	//JSON資料
	data := []byte(` 
{ 
	"message": "Greetings fellow gopher!"
}
`)
	var v greeting                  //建立一個空結構
	err := json.Unmarshal(data, &v) //解析JSON資料和寫入v
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
