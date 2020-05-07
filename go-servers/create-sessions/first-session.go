package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
)

var tpl *template.Template
var userSession = make(map[string]User) //user session mapping table

func init() {
	tpl = template.Must(template.ParseGlob("go-servers/create-sessions/templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8081", nil)
}

type User struct {
	fname string
	lname string
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	var u User
	if err != nil { // if no session id present, create one if post
		fmt.Println(1)
		if req.Method == http.MethodPost {
			fmt.Println(2)
			uuid := strconv.Itoa(int((rand.Float64() * 5) + 5)) //can be implemented using third party
			http.SetCookie(w, &http.Cookie{
				Name:  "session-id",
				Value: uuid,
			})
			u = User{
				req.FormValue("fname"),
				req.FormValue("lname"),
			}
			fmt.Println(3)
			userSession[uuid] = u
		} else {
			//normal page load. Proceed blank
		}
	} else { //session found
		//extract user
		u = userSession[cookie.Value]
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

