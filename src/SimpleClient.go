package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:3000/getjson")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
