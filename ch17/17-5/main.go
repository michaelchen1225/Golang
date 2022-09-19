package main

import "fmt"

func main() {
	helloString := "Hello"
	packtString := "Packt"

	//傳2個引數給Sprintf, 卻只有一個%s
	jointString := fmt.Sprintf("%s %s", helloString, packtString)
	fmt.Println(jointString)
}
