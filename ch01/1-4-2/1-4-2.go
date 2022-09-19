package main

import "fmt"

func main() {
	count := 5 //宣告一個變數並賦予值

	count += 5 //變數加上5,再重新賦值回給同一個變數
	fmt.Println(count)

	count++ //變數增加1
	fmt.Println(count)

	count-- //變數減去1
	fmt.Println(count)

	count -= 5 //變數減去5
	fmt.Println(count)

	name := "John"
	name += " Smith" //也可以自串串接
	fmt.Println("Hello,", name)
}
