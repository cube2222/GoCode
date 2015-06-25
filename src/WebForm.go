package main
import (
	"net/http"
	"io/ioutil"
	"fmt"
)


func main() {
	http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		fmt.Println(req.Form.Get("username"))
		fmt.Println(req.Form.Get("password"))
		fmt.Fprintln(res, "Welcome ", req.Form.Get("username"), ", thank you for logging in!")
	})
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadFile("/home/user/GoCode/src/html/Form.html")
		if err != nil {
			panic(err)
		}
		res.Write(data)
	})

	http.ListenAndServe(":3000", nil)
}
