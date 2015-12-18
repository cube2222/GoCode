package main

import (
	"fmt"
	"net/http"
)

func main() {
	ClientCount := 0

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		ClientCount += 1
		fmt.Fprintf(w, "Hello, %v client!", ClientCount)
	})

	http.ListenAndServe(":80", nil)
}
