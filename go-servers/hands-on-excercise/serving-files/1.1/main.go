package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("go-servers/hands-on-excercise/routing/1.1/index.html"))
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Not found", 404)
	}
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran!")
}

func bb(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "go-servers/hands-on-excercise/routing/1.1/bb.jpg")
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/", foo)
	http.HandleFunc("/bb.jpg", bb)
	http.ListenAndServe(":8080", nil)
}
