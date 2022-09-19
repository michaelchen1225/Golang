package main

import "fmt"

func main() {
	//宣告多重變數並賦予值:
	query, limit, offset := "bat", 10, 0
	//以單行敘述一次更改全部的變數值
	query, limit, offset = "ball", offset, 20

	fmt.Println(query, limit, offset)

}
