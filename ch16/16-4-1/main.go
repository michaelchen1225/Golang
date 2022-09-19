package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	defer close(ch)
	ch <- 1   //在通道傳入1
	i := <-ch //從通道取值並存入i
	fmt.Println(i)
}
