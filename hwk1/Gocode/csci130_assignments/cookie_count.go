package main


import (
	"net/http"
	"io"
	"strconv"
)

func main(){

	http.HandleFunc("/",cookie_count)
	http.ListenAndServe(":8080", nil)
}

func cookie_count(w http.ResponseWriter, req * http.Request){

	cookie, err := req.Cookie("cookie_val")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "cookie_val",
			Value: "0",
		}
	}
	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)

	io.WriteString(w, "value:" + cookie.Value)
}
