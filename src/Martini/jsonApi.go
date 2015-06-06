package main
import (
	"net/http"
	"github.com/go-martini/martini"
	"encoding/json"
)

type Client struct {
	Name string    `json:"name"`
	Id   int    `json:"id"`
}

func main() {
	m := martini.Classic()
	client := Client{"Kuba", 17}
	m.Get("/getjson", func(res http.ResponseWriter, req *http.Request) {
		jts := json.NewEncoder(res)
		jts.Encode(client)
	})
	m.Run()
}
