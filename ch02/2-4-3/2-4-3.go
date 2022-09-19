package main

import "fmt"

func main() {
	config := map[string]string{ //建立map, 元素由一對鍵與值構成
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}

	for key, value := range config { //走訪map並逐次取出鍵/值
		fmt.Println(key, "=", value)
	}
}
