package main

import "fmt"

func greet(ch chan string) {
	msg := <-ch                        //接收訊息1
	ch <- fmt.Sprintf("收到訊息: %s", msg) //傳入訊息2(包含訊息1)
	ch <- "Hello David"                //傳入訊息3
}

func main() {
	ch := make(chan string)
	go greet(ch)

	ch <- "Hello John" //傳入訊息1
	fmt.Println(<-ch)  //接收訊息2
	fmt.Println(<-ch)  //接收訊息2
}
