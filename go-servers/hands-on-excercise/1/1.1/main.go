package main

import (
	"fmt"
	"net/http"
)

func defaultRouteHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "This is a default route")
}

func dogRouteHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "This is a dog route")
}

func myNameRouteHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "This is Tanmay")
}

func main() {
	http.HandleFunc("/", defaultRouteHandler)
	http.HandleFunc("/dog/", dogRouteHandler)
	http.HandleFunc("/me/", myNameRouteHandler)

	http.ListenAndServe(":8080", nil)
}
