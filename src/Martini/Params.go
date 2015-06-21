package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Get("/:name", func(params martini.Params, res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello %v!!!", params["name"])
	})
	m.Run()
}
