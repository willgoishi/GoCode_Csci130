package main


import (
	"net/http"
  "io"
	"fmt"
	/*"strings"*/
)

func main(){

	http.HandleFunc("/",print_out_URL)      //(pattern string, handler func(ResponseWriter, *Request)) parameter
	http.ListenAndServe(":8080", nil)
}

func print_out_URL(w http.ResponseWriter, req * http.Request){

	//g := strings.Split(req.URL.Path,"/")
	//gr:=g[0]

	//io.WriteString(w, "Name:"+ req.URL.Path )
	n:="a"
	value := req.FormValue(n);
        io.WriteString(w, "Name:" + req.URL.Path )
	fmt.Println(value);
	fmt.Println(req.URL.Path)
}
