package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type student struct {
	StudentId     int
	LastName      string
	MiddleInitial string
	FirstName     string
	IsEnrolled    bool
	Courses       []course
}

type course struct {
	Name   string
	Number int
	Hours  int
}

func main() {
	s := student{
		StudentId:  2,
		LastName:   "Washington",
		FirstName:  "Bill",
		IsEnrolled: true,
		Courses: []course{
			{Name: "World Lit", Number: 101, Hours: 3},
			{Name: "Biology", Number: 201, Hours: 4},
			{Name: "Intro to Go", Number: 101, Hours: 4},
		},
	}

	var conn bytes.Buffer                      //模擬通訊用的io.Reader/io.Writer
	encoder := gob.NewEncoder(&conn)           //產生encoder
	if err := encoder.Encode(&s); err != nil { //編碼gob
		fmt.Println("GOB 編碼錯誤:", err)
		os.Exit(1)
	}

	fmt.Printf("%x\n", conn.String()) //把conn的內容用16進位形式印出

	s2 := student{}                             //接收解碼後資料的結構
	decoder := gob.NewDecoder(&conn)            //產生decoder
	if err := decoder.Decode(&s2); err != nil { //解碼gob
		fmt.Println("GOB 解碼錯誤:", err)
		os.Exit(1)
	}

	fmt.Println(s2) //解碼後的資料
}
