package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	StudentId     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"mname,omitempty"`
	FirstName     string   `json:"fname"`
	IsEnrolled    bool     `json:"enrolled"`
	Courses       []course `json:"classes,omitempty"`
}

type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

func main() {
	//第一位學生的資料
	s := student{
		StudentId:     1,
		LastName:      "Williams",
		MiddleInitial: "s",
		FirstName:     "Felicia",
		IsEnrolled:    false,
	}
	//這個學生沒有課程資料, Courses欄位會被略過

	//編碼成JSON時縮排4個空格
	student1, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student1))
	fmt.Println()

	//第二位學生的資料
	s2 := student{
		StudentId:  2,
		LastName:   "Washington",
		FirstName:  "Bill",
		IsEnrolled: true,
	}

	//第二位學生的選課資料
	c := course{Name: "World Lit", Number: 101, Hours: 3}
	s2.Courses = append(s2.Courses, c)
	c = course{Name: "Biology", Number: 201, Hours: 4}
	s2.Courses = append(s2.Courses, c)
	c = course{Name: "Intro to Go", Number: 101, Hours: 4}
	s2.Courses = append(s2.Courses, c)

	student2, err := json.MarshalIndent(s2, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student2))
}
