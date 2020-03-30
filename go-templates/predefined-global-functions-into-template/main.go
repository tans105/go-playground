package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseGlob("go-templates/predefined-global-functions-into-template/*gohtml"))
}

func main() {

	u1 := user{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}

	u2 := user{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		<!DOCTYPE html>
		<html lang="en">
		<head>
		  <meta charset="UTF-8">
		  <title>predefined global functions</title>
		</head>
		<body>
		EXAMPLE #1
		    {Buddha The belief of no beliefs false}
		    {Gandhi Be the change true}
		    { Nobody true}
		EXAMPLE #2
		  EXAMPLE #2 - [{Buddha The belief of no beliefs false} {Gandhi Be the change true} { Nobody true}]
		EXAMPLE #3
		      EXAMPLE #3 - Buddha
		      EXAMPLE #3 - Gandhi
		EXAMPLE #4
		      EXAMPLE #4 - Name: Gandhi
		      EXAMPLE #4 - Motto: Be the change
		      EXAMPLE #4 - Admin: true
		</body>
		</html>
	*/
}
