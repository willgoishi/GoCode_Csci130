package main


import (
	"net/http"
	//"io"
	//"strconv"
	"github.com/nu7hatch/gouuid"
)

func main(){

	http.HandleFunc("/",cookie_count)
	http.ListenAndServe(":8080", nil)
}

func cookie(w http.ResponseWriter, req * http.Request){

	id, _ := uuid.NewV4()

	cookie := &http.Cookie{
		Name:     "session",
		Value:    id.String(),
		HttpOnly: true,
		//Secure: false,
	}
	http.SetCookie(w, cookie)
}
