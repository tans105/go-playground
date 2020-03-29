package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("go-templates/1/templates/*")) //keeps the htmls in a container (respect to main package)
}

func main() {
	err := tpl.Execute(os.Stdout, nil) //bring the templates from the container and shows it to standard output
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil) ////bring the templates from the container and shows it to standard output
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil) ////bring the templates from the container and shows it to standard output
	if err != nil {
		log.Fatalln(err)
	}
}
