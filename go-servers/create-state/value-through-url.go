package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", c)
	http.ListenAndServe(":8081", nil)
}

func c(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Requested String is " +req.FormValue("q"))
}
