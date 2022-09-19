package main

import "fmt"

func genSlices() ([]int, []int, []int) {
	var s1 []int              //用var建立切片
	s2 := make([]int, 10)     //用make建立切片, 指定長度
	s3 := make([]int, 10, 50) //用make建立切片, 指定長度和容量
	return s1, s2, s3
}

func main() {
	s1, s2, s3 := genSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))
}
