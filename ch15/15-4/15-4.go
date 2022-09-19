package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//HTML模板原始字串
var templateStr = `
<html>
  <h1>Customer {{.ID}}</h1>
  {{if .ID }}
    <p>Details:</p>
    <ul>
        {{if .Name}}<li>Name: {{.Name}}</li>{{end}}
        {{if .Surname}}<li>Surname: {{.Surname}}</li>{{end}}
        {{if .Age}}<li>Age: {{.Age}}</li>{{end}}
    </ul>
  {{else}}
    <p>Data not available</p>
  {{end}}
</html>
`

//要來替模板提供資料的結構
type Customer struct {
	ID      int
	Name    string
	Surname string
	Age     int
}

func hello(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query() //取得查詢參數
	customer := Customer{}

	id, ok := vl["id"]
	if ok {
		customer.ID, _ = strconv.Atoi(id[0])
	}

	name, ok := vl["name"]
	if ok {
		customer.Name = name[0]
	}

	surname, ok := vl["surname"]
	if ok {
		customer.Surname = surname[0]
	}

	age, ok := vl["age"]
	if ok {
		customer.Age, _ = strconv.Atoi(age[0])
	}

	//建立名為Exercise15.04的模板, 並填入templateStr模板字串用於解析
	tmpl, _ := template.New("Exercise15.04").Parse(templateStr)
	//使用customer的資料填入模板, 並將結果寫入ResponseWriter (傳給客戶端)
	tmpl.Execute(w, customer)
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
