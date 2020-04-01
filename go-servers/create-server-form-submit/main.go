package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("go-servers/create-server-form-submit/index.html"))
}

func main() {
	var d server
	http.ListenAndServe(":8081", d)
}

type server int

func (s server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Here...")
	err := req.ParseForm()
	if err != nil {
		log.Panic(err)
	}

	tpl.ExecuteTemplate(w,"index.html", req.Form)
}
