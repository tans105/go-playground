package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", d)
	http.ListenAndServe(":8081", nil)
}

func d(w http.ResponseWriter, req *http.Request) {
	tpl = template.Must(template.ParseFiles("go-servers/create-state/templates/form.gohtml"))
	param := req.FormValue("q")
	tpl.ExecuteTemplate(w, "form.gohtml", param)
}
