package main

import "fmt"

var defaultOffset = 10

func main() {
	offset := defaultOffset
	fmt.Println(offset)

	//把offset的值加上defaultOffset的值 , 重新存入offset
	offset = offset + defaultOffset
	fmt.Println(offset)

}
