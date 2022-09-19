package main

import "fmt"

func message() string {
	m := ""
	arr := [4]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]                //將原本的數字平方
		m += fmt.Sprintf("%v: %v\n", i, arr[i]) //用格式化字串將索引和值一一連起來
	}
	return m
}

func main() {
	fmt.Print(message())
}
