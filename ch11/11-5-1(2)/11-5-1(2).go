package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonData := []byte(`
{
    "id": 2,
    "lname": "Washington",
    "fname": "Bill",
    "IsEnrolled": true,
    "grades":[100,76,93,50],
    "class": 
        {
            "coursename": "World Lit",
            "coursenum": 101,
            "coursehours": 3
        }
}	
`)

	if !json.Valid(jsonData) { //先檢查資料是否符合JSON格式
		fmt.Println("JSON 格式不合法:", jsonData)
		os.Exit(1)
	}

	//解碼JSON格式到map
	var v map[string]interface{}
	if err := json.Unmarshal(jsonData, &v); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//走訪map
	for key, value := range v {
		fmt.Printf("%s = %v (%s)\n", key, value, findTypeName(value))
	}
}

//用型別斷言來檢查值的函式
func findTypeName(i interface{}) string {
	switch i.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case float64:
		return "float64"
	case bool:
		return "bool"
	default:
		return fmt.Sprintf("%T", i)
	}
}
