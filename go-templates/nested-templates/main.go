package main

import (
	"log"
	"os"
	"text/template"
)

var tpl  *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("go-templates/nested-templates/templates/*gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
	/*
	<h1>Meaning of Life is 42</h1>
	    This is Tanmay
	    42
	*/
}
