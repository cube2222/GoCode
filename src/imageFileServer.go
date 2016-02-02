package main

import (
	"net/http"
	"os"
	"fmt"
	"io"
)

func main() {
	http.HandleFunc("/myImage", sendImage)
	http.ListenAndServe(":3000", nil)
}

func sendImage(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("C:/temp/image.png")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	io.Copy(w, file)
}
