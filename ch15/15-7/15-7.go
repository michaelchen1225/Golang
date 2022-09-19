package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type WorldTime struct { //對應要回傳的JSON資料結構
	UTC      string `"json:utc"`
	Local    string `"json:local"`
	Timezone string `"json:timezone"`
}

func RestfulService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { //若不是GET方法就回傳HTTP 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	now := time.Now() //取得目前時間

	//讀取參數tz
	vl := r.URL.Query()
	tz, ok := vl["tz"]
	if ok {
		//有tz參數的話, 嘗試用它來設定時區
		location, err := time.LoadLocation(strings.TrimSpace(tz[0]))
		if err != nil { //時區錯誤, 傳回HTTP 404
			w.WriteHeader(http.StatusBadRequest)
			//傳回一個時間欄位為空字串, 時區則帶有錯誤資訊的JSON資料
			jsonBytes, _ := json.Marshal(WorldTime{Timezone: "Invalid timezone name"})
			w.Write(jsonBytes)
			return
		}
		now = now.In(location)
	}

	//取得要傳回的時間和時區字串
	worldTime := WorldTime{}
	worldTime.UTC = now.UTC().Format(time.RFC3339)
	worldTime.Local = now.Format(time.RFC3339)
	worldTime.Timezone = now.Location().String()

	jsonBytes, _ := json.Marshal(worldTime)
	w.Write(jsonBytes)
	// json.NewEncoder(w).Encode(worldTime)
}

func main() {
	http.HandleFunc("/", RestfulService)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
