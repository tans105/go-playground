package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("go-servers/hands-on-excercise/1/1.2/index.gohtml"))
}

func defaultRouteHandler(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", "Default")
	if err != nil {
		log.Fatalln(err)
	}

}

func dogRouteHandler(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", "Dog")
	if err != nil {
		log.Fatalln(err)
	}
}

func myNameRouteHandler(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", "Tanmay")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(defaultRouteHandler))
	http.Handle("/dog/", http.HandlerFunc(dogRouteHandler))
	http.Handle("/me/", http.HandlerFunc(myNameRouteHandler))

	http.ListenAndServe(":8080", nil)
}
