package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

const (
	firstName = iota //CSV欄位索引
	lastName
	age
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	header := true // 標頭開關
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if header {
			header = false
			continue //跳過第一行(標頭)
		}
		fmt.Println("--------------------")
		fmt.Println("First name:", record[firstName])
		fmt.Println("Last name :", record[lastName])
		fmt.Println("Age       :", record[age])
	}
}
