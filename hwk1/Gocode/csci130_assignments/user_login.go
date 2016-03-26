package main

import (
	"log"
	//"text/template"
	"net/http"
	//"github.com/nu7hatch/gouuid"
	//"fmt"
	"encoding/json"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha256"
	"io/ioutil"
	"io"
)

type user struct{
	name string
	age string
	uuid string
	HMAC string
}

var loginFile string
var viewFile string
func init() {
	temp, _ := ioutil.ReadFile("project_html.html")
	temp1, _ := ioutil.ReadFile("project_html.html")
	viewFile = string(temp1)
	loginFile = string(temp)
}

//HMAC
func HMACFunction(data string) string {
	hmac := hmac.New(sha256.New, []byte(data))
	return string(hmac.Sum(nil))
}
func undoJSON(cookie *http.Cookie) (user, bool) {
	decode, _ := base64.StdEncoding.DecodeString(cookie.Value)
	var jsonVal user
	json.Unmarshal(decode, &jsonVal)
	if hmac.Equal([]byte(jsonVal.HMAC), []byte(HMACFunction(jsonVal.uuid+jsonVal.name+jsonVal.age))) {
		return jsonVal, true
	}
	return jsonVal, false
}

func reJSON(jsonValue user) string {
	encode, _ := json.Marshal(jsonValue)
	return base64.StdEncoding.EncodeToString(encode)
}

func make_cookie(cookie *http.Cookie, req *http.Request) string{
	jsonVal, _ := undoJSON(cookie)
	jsonVal.name = req.FormValue("name")
	jsonVal.age = req.FormValue("age")
	jsonVal.HMAC = req.FormValue("HMAC")
	return reJSON(jsonVal)
}

func main(){

	http.HandleFunc("/",Loginserving)

	error := http.ListenAndServe(":8080", nil)

	if error != nil{
		log.Fatal("Error",error)
	}


}
func Loginserving(result http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("logged-in")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "logged-in",
			Value: "0",
			//Secure: true,
			HttpOnly: true,
		}
	}

	// check log in
	if req.Method == "POST" {
		password := req.FormValue("password")
		if password == "pleb" {
			cookie = &http.Cookie{
				Name:  "logged-in",
				Value: "1",
				//Secure: true,
				HttpOnly: true,
			}


		}
	}
	if req.URL.Path == "/logout" {
		cookie = &http.Cookie{
			Name:   "logged-in",
			Value:  "0",

			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(result, cookie)
		http.Redirect(result, req, "/", 303)
		return
	}

	http.SetCookie(result, cookie)
	var html string
	if cookie.Value == "0" {
		html = `
			<!DOCTYPE html>
			<html lang="en">
			  	<head>
					<meta charset="UTF-8">
				</head>
				<body>
					<h1>Login</h1>
			             <form method="POST">
					<h3>password here</h3>
				<input type="text" name="password">
				  <br>
				<input type="submit">
				     </form>
				</body>
			</html>`
	}
	if cookie.Value == "1" {
		html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<h1><a href="/logout">LOG OUT</a></h1>
			</body>
			</html>`
		Loginserving(result,req)
		make_cookie(cookie, req)
	}

	io.WriteString(result, html)
}



/*func serves(result http.ResponseWriter, req *http.Request){
	temp:=template.New("project_html.html")

	u:=user{
		name : "will",
		age: "20",
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

	if req.Method == "POST" {
		cookie.Value = make_cookie(cookie, req)
	}
	obj, status := undoJSON(cookie)
	vobj, _ := undoJSON(cookie)

	if status {
		t, _ := template.New("name").Parse("file")
		t.Execute(result, obj)
	}

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


	t, _ := template.New("Name").Parse(loginFile)
	v, _ := template.New("Name").Parse(viewFile)
	t.Execute(result, obj)
	v.Execute(result, vobj)
}*/
