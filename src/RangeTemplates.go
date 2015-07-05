package main
import (
	"os"
	"io/ioutil"
	"html/template"
	"net/http"
)

type Information struct {
	Names [5]string
}

func main() {
	var myInfo Information
	myInfo.Names[0] = "Ben"
	myInfo.Names[1] = "Tom"
	myInfo.Names[2] = "Jim"
	file, err := os.Open("C:/Users/User/Documents/GoCode/src/html/RangeTemplate.html")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)
	tmpl, err := template.New("test").Parse(string(data))

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		tmpl.Execute(res, myInfo)
	})
	http.ListenAndServe(":3000", nil)
}
