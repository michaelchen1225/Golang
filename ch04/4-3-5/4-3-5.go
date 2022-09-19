package main

import "fmt"

func message() string {
	arr := [4]string{"ready", "Get", "Go", "to"}
	arr[1] = "It's" //修改元素值
	arr[0] = "time"
	//輸出修改後的元素值
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}

func main() {
	fmt.Print(message())
}
