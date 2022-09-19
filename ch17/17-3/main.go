package main

import "fmt"

func main() {
	finished := make(chan bool)
	names := []string{"Packt"}

	//goruoutine嘗試在names加入值
	go func() {
		names = append(names, "Electric")
		names = append(names, "Boogaloo")
		finished <- true
	}()
	//但同時又被main()嘗試讀取names
	for _, name := range names {
		fmt.Println(name)
	}
	<-finished
}
