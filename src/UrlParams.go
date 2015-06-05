package main
import (
	"net/http"
	"net/url"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		values, _ := url.ParseQuery(req.URL.RawQuery)
		fmt.Fprintln(res, values["variable"][0] + " " + values["name"][0])
	})
	http.ListenAndServe(":3000", nil)
}
