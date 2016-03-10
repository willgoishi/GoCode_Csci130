package main

import (
	"log"
	"text/template"
	"net/http"
	"github.com/nu7hatch/gouuid"
	"fmt"
)

type student struct{
	name string
	person_age int
	major string
	standing string
	study_level string
}

func serve(result http.ResponseWriter, req *http.Request){
	temp:=template.New("project_html.html")

	s:=student{
		name : "will",
		person_age : 20,
		major : "csci",
		standing : "Senior",
	}

	temp,err := temp.ParseFiles("project_html.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(result, s)

        //Cookie Section
	cookie, err := req.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(result, cookie)
	}
	temp.Execute(result, nil)
	fmt.Println(cookie)

}


func main(){

	http.HandleFunc("/",serve)

	error := http.ListenAndServe(":9000", nil)

	if error != nil{
		log.Fatal("Error",error)
	}


}
