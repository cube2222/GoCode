package main
import (
	"net/http"
	"html/template"
	"fmt"
	"os"
	"io/ioutil"
)

type Human struct {
	Name string
	Age  int
}

func main() {
	f, err := os.Open("/home/user/GoCode/src/html/SimpleTemplate.html")
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

	testSubject := Human{"Kuba", 17}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		err := tmpl.Execute(res, testSubject)
		if err != nil {
			fmt.Println(err)
		}
	})
	http.ListenAndServe(":3000", nil)
}
