package main

import (
	"log"
	"text/template"
	"net/http"
	"github.com/nu7hatch/gouuid"
	"fmt"
	"encoding/json"
	"encoding/base64"
)

type user struct{
	name string
	age int
}

func serve(result http.ResponseWriter, req *http.Request){
	temp:=template.New("project_html.html")

	u:=user{
		name : "will",
		age: 20,
	}

	temp,err := temp.ParseFiles("project_html.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(result, u)

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
	//temp.Execute(result, nil)
	fmt.Println(cookie)

        //Form gathers user name and age and stores it into the cookie
	userName := req.FormValue("name")
	userAge := req.FormValue("age")

	cookie.Value = "|" + userName + "|" + userAge

	//Here is where the values are collected and set to the new user variable
	client := user{
		name:userName ,
		age:userAge ,
	}

        //marshall variable of type user to json
        abc,_:=json.Marshal(client)

	//encode from json to base 64
	str := base64.URLEncoding.EncodeToString(abc)
	
	//store that into the cookie
        cookie.Value = str
	
	http.SetCookie(result, cookie)
        fmt.Println()
}


func main(){

	http.HandleFunc("/",serve)

	error := http.ListenAndServe(":9000", nil)

	if error != nil{
		log.Fatal("Error",error)
	}


}
