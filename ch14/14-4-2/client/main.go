package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func postFileAndReturnResponse(filename string) string {
	fileDataBuffer := bytes.Buffer{}                  //建立一個buffer
	mpWritter := multipart.NewWriter(&fileDataBuffer) //建立multipart.Writer

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//用multipart.Writer 建立準備傳送的MIME檔案
	formFile, err := mpWritter.CreateFormFile("myFile", file.Name())
	if err != nil {
		log.Fatal(err)
	}

	//將原始檔案的內容拷貝到MIME
	if _, err := io.Copy(formFile, file); err != nil {
		log.Fatal(err)
	}
	mpWritter.Close()

	//用POST請求送出MIME檔案並讀取回應
	//使用multipart.Writer來指定標頭內的內容類型為 multipart/form-data
	r, err := http.Post("http://localhost:8080", mpWritter.FormDataContentType(), &fileDataBuffer)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	data := postFileAndReturnResponse("./test.txt")
	fmt.Println(data)
}
