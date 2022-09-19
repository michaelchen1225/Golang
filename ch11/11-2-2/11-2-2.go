package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type greeting struct {
	SomeMessage string `json:"message"`
}

func main() {
	//JSON資料
	data := []byte(` 
{ 
	"message": "Greetings fellow gopher!"
}
`)
	if !json.Valid(data) { //檢查JSON格式是否正確
		fmt.Printf("JSON 格式無效: %s", data)
		os.Exit(1)
	}

	v := greeting{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
