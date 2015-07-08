package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Human struct {
	Name string
	Age  int
}

func main() {
	f, err := os.Open("src/html/SimpleTemplate.html")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	tmpl, err := template.New("test").Parse(string(data))
	if err != nil {
		fmt.Println(err)
	}

	testSubject := Human{"Kuba", 17}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		err := tmpl.Execute(res, testSubject)
		if err != nil {
			fmt.Println(err)
		}
	})
	http.ListenAndServe(":3000", nil)
}
