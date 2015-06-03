package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("http://localhost:3000")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
