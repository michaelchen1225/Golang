package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "user:1234@tcp(localhost:3306)/mysqldb?charset=utf8")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("sql.DB 結構已建立")

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("資料庫連線成功")

	//查詢資料表, 傳回sql.Rows
	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		panic(err)
	}
	defer rows.Close() //在程式結束時關閉Rows
	fmt.Println("資料表查詢成功, 列出 employee 內容...")

	for rows.Next() { //走走訪Rows
		e := employee{}
		err := rows.Scan(&e.id, &e.name) //讀出一筆資料
		if err != nil {
			panic(err)
		}
		fmt.Println(e.id, e.name) //印出資料
	}
	err = rows.Err() //檢查Rows有無遭遇其他錯誤
	if err != nil {
		panic(err)
	}
}
