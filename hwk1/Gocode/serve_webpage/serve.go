
package main

import (
        "html/template"
        "net/http"
	"log"
)

func serve(result http.ResponseWriter, req *http.Request){
         t := template.New("index.gohtml")
	 t,err := template.ParseFiles("index.gohtml")

	if err != nil {
		log.Fatal("Error:",nil)
	}
	t.Execute(result ,nil)

}
func main(){

	http.Handle("/css/", http.StripPrefix("/css/",http.FileServer(http.Dir("css"))))
        http.Handle("/image/", http.StripPrefix("/image/",http.FileServer(http.Dir("image"))))
	http.HandleFunc("/",serve)
	error := http.ListenAndServe(":9000", nil)

	if error != nil{
             log.Fatal("Error",error)
	}
}
