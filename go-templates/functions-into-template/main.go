package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("go-templates/functions-into-template/templates/*html"))
}

func firstThree(s string) string {
	return strings.TrimSpace(s[:3])
}

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "functions.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

	/*
	<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <title>functions</title>
	</head>
	<body>
	    BUDDHA
	    GANDHI
	    MARTIN LUTHER KING

	    Bud
	    Gan
	    Mar
	</body>
	</html>
	*/
}
