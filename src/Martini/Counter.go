package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	m := martini.Classic()
	counter := 0
	m.Get("/", func(res http.ResponseWriter, req *http.Request) {
		counter += 1
		fmt.Fprintf(res, "Hello %v client!!!", counter)
	})
	m.Run()
}
