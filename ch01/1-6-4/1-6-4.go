package main

import "fmt"

func add5Value(count int) { //傳入值
	count += 5
	fmt.Println("add5Value     :", count)
}

func add5Point(count *int) { //傳入指標
	*count += 5
	fmt.Println("add5Point     :", *count)
}

func main() {
	var count int
	add5Value(count)
	fmt.Println("add5Value post:", count)
	add5Point(&count)
	fmt.Println("add5Point post:", count)
}
