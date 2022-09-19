package main

import (
	"html/template"
	"log"
	"net/http"
)

type Visitor struct { //用來整理使用者以表單送出的資料
	Name    string
	Surname string
	Age     string
}

type Hello struct { //請求處理器
	tmpl *template.Template //紀錄模板物件的屬性
}

//請求處理器方法(請求/hello路徑時)
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	visitor := Visitor{}

	//檢查是不是POST請求, 不是就回傳HTTP狀態碼405 (method not allowed)
	if r.Method != http.MethodPost { // 也可以寫成 r.Method != "POST"
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		//沒有表單或解讀錯誤的話, 回傳HTTP狀態碼400(bad request)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//從表單讀取值
	visitor.Name = r.Form.Get("name")
	visitor.Surname = r.Form.Get("surname")
	visitor.Age = r.Form.Get("age")

	h.tmpl.Execute(w, visitor)
}

//用來設定並傳回請求處理器的函式
func NewHello(tmplPath string) (*Hello, error) {
	//設定請求處理器要使用的模板
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return nil, err
	}

	return &Hello{tmpl}, nil
}

func main() {
	hello, err := NewHello("./result.html") //取得請求處理器
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/hello", hello) //要hello處理 /hello(表單) 路徑的請求

	//伺服器根路徑則指向form.html表單畫面
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./form.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
