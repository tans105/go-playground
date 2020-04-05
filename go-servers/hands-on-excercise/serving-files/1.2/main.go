package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("go-servers/hands-on-excercise/serving-files/1.2"))))
}
