package main

import (
	"fmt"
	"net/http"
)

func main() {
	var d server
	//http.ListenAndServe(port string, Handler interface)
	//any type which implements ServeHTTP(w http.ResponseWriter, req *http.Request) will be implementing handler interface
	http.ListenAndServe(":8081", d)
}

type server int

func (s server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Woops this server works ! ")
}
