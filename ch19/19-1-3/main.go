package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name    string  `des:"userName"` //欄位帶有標籤
	Age     int     `des:"userAge"`
	Balance float64 `des:"bankBalance"`
	Member  bool    `des:"isMember"`
}

//用reflect檢視結構內容
func PrintStruct(s interface{}) {
	sT := reflect.TypeOf(s)  //取得reflect.Type
	sV := reflect.ValueOf(s) //取得reflect.Value

	//印出結構型別名稱和其基礎型別名稱
	fmt.Printf("type %s %v {\n", sT.Name(), sT.Kind().String())

	//走訪結構欄位
	for i := 0; i < sT.NumField(); i++ {
		field := sT.Field(i) //取得第i個欄位的型別 (reflect.Type)
		value := sV.Field(i) //取得第i個欄位的值 (reflect.Value)

		//印出欄位名稱/型別/值 以及標籤des的字串
		fmt.Printf("\t%s\t%s\t= %v\t(description: %s)\n",
			field.Name, field.Type.String(),
			value.Interface(), field.Tag.Get("des"))
	}

	fmt.Println("}")
}

func main() {
	u1 := User{
		Name:    "Tracy",
		Age:     51,
		Balance: 98.43,
		Member:  true,
	}

	PrintStruct(u1) //用reflect印出u1內容

	//透過u1的指標用reflect指名欄位名稱, 以便更改欄位值
	v := reflect.ValueOf(&u1)
	v.Elem().FieldByName("Name").SetString("Grace")
	v.Elem().FieldByName("Age").SetInt(45)
	v.Elem().FieldByName("Balance").SetFloat(56.97)
	v.Elem().FieldByName("Member").SetBool(false)

	PrintStruct(u1) //再次印出u1內容
}
