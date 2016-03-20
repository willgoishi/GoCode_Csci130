package main


import (
	"net/http"
	"io"
	//"fmt"
	/*"strings"*/
	"io/ioutil"
	"os"
	"path/filepath"
)

func main(){

	http.HandleFunc("/",submit_form)
	http.ListenAndServe(":8080", nil)
}

func submit_form(w http.ResponseWriter, req * http.Request){

	if req.Method == "POST" {
		file, _, err := req.FormFile("input")
		if err !=nil{
			http.Error(w, err.Error(),500)
			return
		}
		defer file.Close()
		src := io.LimitReader(file,404)
		dst, err := os.Create(filepath.Join(".","input.txt"))
		if err != nil{
			http.Error(w,err.Error(),500)
			return
		}
		defer dst.Close()
		io.Copy(dst, src)
		dat, err := ioutil.ReadFile("input.txt")
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, string(dat))
	}
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `
                           <form method="POST" enctype="multipart/form-data">
                           <input type="file" name="input">
                           <input type="submit">
                           </form>`)
}
