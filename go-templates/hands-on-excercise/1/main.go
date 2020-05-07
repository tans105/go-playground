package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number string
	Name   string
	Units  string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcaYear string
	Fall    semester
	Spring  semester
	Summer  semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("go-templates/hands-on-excercise/1/tpl.gohtml"))
}

func main() {
	years := []year{
		year{
			AcaYear: "2020-2021",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CSCI-40", "Introduction to Programming in Go", "4"},
					course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
		year{
			AcaYear: "2021-2022",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CSCI-40", "Introduction to Programming in Go", "4"},
					course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, years)
	if err != nil {
		log.Fatalln(err)
	}

	/*
	<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <title>Document</title>
	</head>
	<body>
	  <li>2020-2021</li>
	  <li>Fall</li>
	      {CSCI-40 Introduction to Programming in Go 4}
	      {CSCI-130 Introduction to Web Programming with Go 4}
	      {CSCI-140 Mobile Apps Using Go 4}
	  <li>Spring</li>
	      {CSCI-50 Advanced Go 5}
	      {CSCI-190 Advanced Web Programming with Go 5}
	      {CSCI-191 Advanced Mobile Apps With Go 5}

	  <li>2021-2022</li>
	  <li>Fall</li>
	      {CSCI-40 Introduction to Programming in Go 4}
	      {CSCI-130 Introduction to Web Programming with Go 4}
	      {CSCI-140 Mobile Apps Using Go 4}
	  <li>Spring</li>
	      {CSCI-50 Advanced Go 5}
	      {CSCI-190 Advanced Web Programming with Go 5}
	      {CSCI-191 Advanced Mobile Apps With Go 5}
	</body>
	</html>
	*/
}
