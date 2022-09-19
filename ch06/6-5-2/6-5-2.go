package main

import (
	"errors"
	"fmt"
)

func main() {
	msg := "good-bye"
	message(msg)
	fmt.Println("這行不會印出")
}

func message(msg string) {
	if msg == "good-bye" {
		panic(errors.New("出事了"))
	}
}
