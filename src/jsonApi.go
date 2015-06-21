package main

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func main() {
	client := Client{"Kuba", 17}
	http.HandleFunc("/getjson", func(res http.ResponseWriter, req *http.Request) {
		jts := json.NewEncoder(res)
		jts.Encode(client)
	})
	http.ListenAndServe(":3000", nil)
}
