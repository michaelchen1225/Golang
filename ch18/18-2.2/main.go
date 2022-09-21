package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "mysecretpassword"
	fmt.Println("密碼明碼  :", password)

	//用bcrypt將密碼轉成雜湊值
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("密碼雜湊值:", string(hash))

	//測試輸入的新密碼是否符合
	testString := "mysecretpassword"
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(testString))
	if err != nil {
		fmt.Println("密碼不符")
	} else {
		fmt.Println("密碼相符")
	}
}
