package main


import (
	"fmt"
	"net/http"
)

func serves(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res,
		`<!DOCTYPE html>
			<html lang="en">
			  <head>
			    <meta charset="UTF-8">
			    <title>TLS Web Page</title>
			  </head>
			</html>`)
}

func redTLS(res http.ResponseWriter, req *http.Request){
	http.Redirect(res, req, req.RequestURI, http.StatusMovedPermanently)
}

func main(){

	http.HandleFunc("/",serves)
	http.ListenAndServe(":8080", nil)
	go http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}
