package main
import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{name}/{age}", func (w http.ResponseWriter, req *http.Request) {
		variables := mux.Vars(req)
		fmt.Fprintf(w, "Hello %v aged %v, nice to meet you!", variables["name"], variables["age"])
	})
	http.ListenAndServe(":3000", r)
}
