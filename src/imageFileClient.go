package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

func main() {
	r, err := http.Get("http://localhost:3000/myImage")
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create("C:/temp/myDownloadedImage.png")
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(file, r.Body)
}
