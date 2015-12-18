package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Human struct {
	Name string
	Age  string
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
	//tmpl, err := template.New("test").Parse("Hello {{.Name}} aged {{.Age}}. Nice to meet you!!!")
	tmpl, err := template.New("test").Parse(string(data))
	if err != nil {
		fmt.Println(err)
	}

	m := martini.Classic()
	m.Get("/:name/:age", func(params martini.Params, res http.ResponseWriter, req *http.Request) {
		testSubject := Human{params["name"], params["age"]}
		err = tmpl.Execute(res, testSubject)
		if err != nil {
			fmt.Println(err)
		}
	})
	m.Run()
}
