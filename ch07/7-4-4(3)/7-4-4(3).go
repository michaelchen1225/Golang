package main

import (
	"fmt"
)

type person struct { //自訂結構
	lastName  string
	age       int
	isMarried bool
}

type animal struct { //自訂結構
	name     string
	category string
}

type record struct { //用來整理map元素的結構
	key       string
	data      interface{}
	valueType string
}

func main() {
	m := make(map[string]interface{}) //建立並初始化一個map

	//在map內加入不同型別的多筆資料
	m["person"] = person{lastName: "Doe", isMarried: false, age: 19}
	m["firstname"] = "Smith"
	m["age"] = 54
	m["isMarried"] = true
	m["animal"] = animal{name: "oreo", category: "cat"}

	//解析map的每個元素, 轉換成record結構和范進一個切片
	rs := []record{}
	for k, v := range m {
		rs = append(rs, newRecord(k, v))
	}

	//印出record結構切片的內容
	for _, v := range rs {
		fmt.Println("Key: ", v.key)
		fmt.Println("Data: ", v.data)
		fmt.Println("Type: ", v.valueType)
		fmt.Println()
	}
}

//處理map資料並輸出程結構
func newRecord(key string, i interface{}) record {
	r := record{}
	r.key = key
	switch v := i.(type) {
	case int:
		r.valueType = "int"
		r.data = v
		return r
	case bool:
		r.valueType = "bool"
		r.data = v
		return r
	case string:
		r.valueType = "string"
		r.data = v
		return r
	case person:
		r.valueType = "person"
		r.data = v
		return r
	default:
		r.valueType = "unknown"
		r.data = v
		return r
	}
}
