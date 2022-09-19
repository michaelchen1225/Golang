package main

import "fmt"

func message() string {
	arr := [...]string{
		"ready",
		"Get",
		"Go",
		"to",
	}
	//用fmt.Sprintln()回傳格式化自串
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}

func main() {
	fmt.Print(message()) //印出格式化字串
}
