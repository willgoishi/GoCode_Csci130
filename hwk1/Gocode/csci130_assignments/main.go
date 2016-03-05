package main


import (
	"net/http"
        "io"
)
func main(){

	http.HandleFunc("/",print_out_URL)      //(pattern string, handler func(ResponseWriter, *Request)) parameter
	http.ListenAndServe(":8080", nil)
}

func print_out_URL(w http.ResponseWriter, req * http.Request){
	io.WriteString(w, "Here is the URL Path"+req.URL.Path)
}
