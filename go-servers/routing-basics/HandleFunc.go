package main

import (
	"io"
	"net/http"
)

func main() {
	//HandleFunc will take in pattern and a function whereas Handle will take in path and Handler interface
	http.HandleFunc("/dog/", d) // takes in anything with /dog eg. /dog/there/is/something/i/dont/care
	http.HandleFunc("/cat", c) // takes in /cat only

	http.ListenAndServe(":8081", nil) //implies we are passing the default Mux
}


func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Its a cat!")
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Its a dog!")
}
