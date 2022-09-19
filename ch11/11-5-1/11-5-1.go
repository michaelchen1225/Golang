package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//原始資料
	jsonData := []byte(`{"checkNum":123,"amount":200,"category":["gift","clothing"]}`)
	//定義map
	var v map[string]interface{}

	//將JSON資料解碼到map
	json.Unmarshal(jsonData, &v)

	//印出map的內容
	fmt.Println(v)
	for key, value := range v {
		fmt.Println(key, "=", value)
	}
}
