import (
	"net/http"
        "io"
	//"fmt"
	/*"strings"*/
)

func main(){

	http.HandleFunc("/",print_out_URL)      //(pattern string, handler func(ResponseWriter, *Request)) parameter
	http.ListenAndServe(":8080", nil)
}

func print_out_URL(w http.ResponseWriter, req * http.Request){

	val := req.FormValue("name")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<form method="POST">
			     Enter your name:
		           <input type="text" name="name">
		           <input type="submit">
                           </form>`+val)


}
