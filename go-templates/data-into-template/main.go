package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("go-templates/data-into-template/templates/*html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "primitive.gohtml", 42)
	if err != nil {
		log.Fatalln()
	}

	/*
		<html>
		<body>
		<h1>Meaning of life is 42</h1>

		<h1>Meaning of life is 42</h1>
		</body>
		</html>
	*/

	sages := []string{"A", "B", "C", "D"}
	err = tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", sages)
	if err != nil {
		log.Fatalln()
	}

	/*
		<html>
		<body>
		  <li>0-A</li>
		  <li>1-B</li>
		  <li>2-C</li>
		  <li>3-D</li>
		</body>
		</html>
	*/

	myMap := map[string]string{
		"India":    "A",
		"Pakistan": "B",
		"USA":      "C",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "maps.gohtml", myMap)
	if err != nil {
		log.Fatalln()
	}

	/*
		<html>
		<body>
		  <li>India-A</li>
		  <li>Pakistan-B</li>
		  <li>USA-C</li>
		</body>
		</html>
	*/

	type city struct {
		Name string
		Id   int
	}

	err = tpl.ExecuteTemplate(os.Stdout, "structs.gohtml", city{"Kota", 1})
	if err != nil {
		log.Fatalln()
	}

	/*
		<html>
		<body>
		Kota - 1
		</body>
		</html>
	*/
}
