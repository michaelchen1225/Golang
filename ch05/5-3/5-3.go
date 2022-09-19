package main

import (
	"fmt"
)

func main() {
	i := []int{5, 10, 15}
	nums(i...)
}

func nums(i ...int) {
	fmt.Println(i)
}
