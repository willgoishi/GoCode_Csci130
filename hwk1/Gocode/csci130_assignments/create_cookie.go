package main


import (
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"net/http"
)

func main(){

	http.HandleFunc("/",cookie)
	http.ListenAndServe(":8080", nil)
}

func cookie(w http.ResponseWriter, req * http.Request){

	cookie, err := req.Cookie("session-id")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:     "session-id",
			Value:    "",
			HttpOnly: true,
			//Secure:true,
		}
	}
	if req.FormValue("name") != "" {
		cookie.Value = cookie.Value + `name= ` + SHA("name")
	}

	http.SetCookie(w, cookie)
	io.WriteString(w, `<!DOCTYPE html>
		<html>
		<body>
		       <form method="POST">
		           <label for="Name">What is your name? -> </label>
		           <input type="text" name="name">
		           <input type="submit" value="Enter!">
		       </form>
		       <a href="/auth">Valid HMAC</a>
		</body>
		</html>`)

}

func SHA(data string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func authentication(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("cookie")
	if err != nil {
		http.Redirect(res, req, "/", 303)
		return
	}
	if cookie.Value == "" {
		http.Redirect(res, req, "/", 303)
		return
	}

}
