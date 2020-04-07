package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var tplt *template.Template

func main() {
	http.HandleFunc("/", e)
	http.ListenAndServe(":8081", nil)
}

func e(w http.ResponseWriter, req *http.Request) {
	tplt = template.Must(template.ParseFiles("go-servers/create-state/templates/fileForm.gohtml"))
	var str string

	if req.Method == http.MethodPost {
		file, _, err := req.FormFile("q")
		if err != nil {
			log.Fatalln(err)
		}

		defer file.Close()

		byteStream, err := ioutil.ReadAll(file)

		if err != nil {
			log.Fatalln(err)
		}

		str = string(byteStream)
	}

	tplt.ExecuteTemplate(w, "fileForm.gohtml", str)
}
