package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("go-servers/hands-on-excercise/serving-files/1.3/templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("go-servers/hands-on-excercise/serving-files/1.3/public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", dogs)
	http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
