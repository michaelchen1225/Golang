package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	//產生0-999之間的亂數
	r, _ := rand.Int(rand.Reader, big.NewInt(1000))
	fmt.Println(r)
}
