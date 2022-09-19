package main

import (
	"errors"
	"fmt"
)

func main() {
	ErrBadData := errors.New("some bad data")
	fmt.Printf("ErrBadData Type : %T", ErrBadData)
}
