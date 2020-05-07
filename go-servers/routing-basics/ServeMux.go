package main

import (
	"io"
	"net/http"
)

type hotdog1 int

func (d hotdog1) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat1 int

func (c hotcat1) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux() //serveMux implements Handler interface
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux) //No Handler means its using default
}
