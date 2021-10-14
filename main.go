package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) { //Function that handle a particular web request
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc) //Function giving path and calling handlerFunc
	http.ListenAndServe(":8080", nil) //Start the Web Server
}
