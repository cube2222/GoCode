package main

import (
	"fmt"
	"net/http"
	"net/url"
)

/*
Example usage:
http://localhost:3000/?variable=4&name=HelloWorld
 */

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		values, _ := url.ParseQuery(req.URL.RawQuery)
		fmt.Fprintln(res, values["variable"][0]+" "+values["name"][0])
	})
	http.ListenAndServe(":3000", nil)
}
