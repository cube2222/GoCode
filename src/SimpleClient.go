package main
import (
	"net/http"
	"fmt"
	"io/ioutil"
)



func main() {
	res, err := http.Get("http://localhost:3000/getjson")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
