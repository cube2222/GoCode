package main

import (
	"github.com/go-martini/martini"
	"net/http"
	"fmt"
)

func main() {
	m := martini.Classic()
	m.Use(Auth)
	m.Get("/:name", func(params martini.Params, res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello %v!!!", params["name"])
	})
	m.Run()
}

func Auth(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("API-KEY") != "password123" {
		http.Error(res, "Nope", 401)
	}
}