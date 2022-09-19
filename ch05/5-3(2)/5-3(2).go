package main

import (
	"fmt"
)

func main() {
	i := []int{5, 10, 15}
	fmt.Println(sum(5, 4))
	fmt.Println(sum(i...))
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums { //走訪數量不定參數的切片和加總
		total += num
	}
	return total
}
