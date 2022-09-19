package main

import "fmt"

type user struct { //自訂結構型別
	name    string //定義欄位名稱與型別
	age     int
	balance float64
	member  bool
}

func getUsers() []user {
	u1 := user{
		name:    "Tracy", //搭配欄位名稱賦值法 (不必照順序)
		age:     51,
		balance: 98.43,
		member:  true,
	}
	u2 := user{
		age:  19,
		name: "Nick",
		//其餘沒有賦值的欄位會是零值
	}
	u3 := user{
		"Bob", //不使用欄位名稱的賦值法(必須照順序 , 資料不能有缺)
		25,
		0,
		false,
	}
	var u4 user
	u4.name = "Sue" //透過 "結構.欄位 = 值" 賦值
	u4.age = 31
	u4.member = true
	u4.balance = 17.09

	return []user{u1, u2, u3, u4}
}

func main() {
	users := getUsers()
	for i := 0; i < len(users); i++ {
		fmt.Printf("%v: %#v\n", i, users[i])
	}
}
