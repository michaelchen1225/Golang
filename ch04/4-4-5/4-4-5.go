package main

import "fmt"

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1    //直接複製切片 (指向同一個陣列)
	s3 := s1[:] //從切片建立相同長度的切片(指向同一個陣列)
	s1[3] = 99
	return s1[3], s2[3], s3[3]
}

func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s1 = append(s1, 6) //家入新值, 超過s1的容量, 使s1置換陣列
	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10) //將s1容量設為10
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, 6) //家入新值, 不超過s1容量, s1不會置換陣列
	s1[3] = 99         //更動仍會反映在s2
	return s1[3], s2[3]
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, []int{10: 11}...) //家入新值(索引10), 超過s1容量
	s1[3] = 99
	return s1[3], s2[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1)) //用make建立與s1同長的切片
	copied := copy(s2, s1)     //用copy複製切片並傳回複製的數量
	s1[3] = 99
	return s1[3], s2[3], copied
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...) //將s1的元素附加到新的切片
	s1[3] = 99
	return s1[3], s2[3]
}

func main() {
	l1, l2, l3 := linked()
	fmt.Println("有連結               :", l1, l2, l3)
	nl1, nl2 := noLink()
	fmt.Println("無連結               :", nl1, nl2)
	cl1, cl2 := capLinked()
	fmt.Println("有設容量, 有連結     :", cl1, cl2)
	cnl1, cnl2 := capNoLink()
	fmt.Println("有設容量, 無連結     :", cnl1, cnl2)
	copy1, copy2, copied := copyNoLink()
	fmt.Print("使用 copy(), 無連結  : ", copy1, copy2)
	fmt.Printf(" (複製了 %v 個元素)\n", copied)
	a1, a2 := appendNoLink()
	fmt.Println("使用 append(), 無連結:", a1, a2)
}
