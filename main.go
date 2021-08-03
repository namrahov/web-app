package main

import (
	"github.com/namrahov/web-app/handlers"
	"net/http"
)

var portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.Yeni)

	_ = http.ListenAndServe(portNumber, nil)

}
