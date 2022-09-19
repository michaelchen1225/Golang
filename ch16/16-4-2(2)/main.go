package main

import "fmt"

func push(from, to int, in chan bool, out chan int) {
	for i := from; i <= to; i++ {
		<-in     //等待請求(值不重要)
		out <- i //傳回一個值
	}
}

func main() {
	s1 := 0
	out := make(chan int, 100) //用來接收值的通道
	in := make(chan bool, 100) //用來送出請求的通道

	go push(1, 25, in, out)
	go push(26, 50, in, out)
	go push(51, 75, in, out)
	go push(76, 100, in, out)

	for c := 0; c < 100; c++ {
		in <- true //送出一個請求
		i := <-out //接收一個數字
		fmt.Println(i)
		s1 += i
	}

	fmt.Println("Result:", s1)
}
