package main
import (
	"net/http"
	"time"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("TestCookie")
		if err != nil {
			if err == http.ErrNoCookie {
				cookie = new(http.Cookie)
				cookie.Name = "TestCookie"
				cookie.Value = "123123123"
				cookie.Domain = "http://localhost:3000"
				cookie.Expires = time.Date(2015, time.July, 10, 1, 1, 1, 1, time.UTC)
				cookie.MaxAge = 180
				http.SetCookie(res, cookie)
				fmt.Fprintf(res, "Sending you a cookie!!!")
			} else {
				fmt.Fprintf(res, "Undefined error")
			}
		} else {
			fmt.Fprintf(res, "I found your cookie!!!")
		}
	})
	http.ListenAndServe(":3000", nil)
}
