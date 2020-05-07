package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", write)
	http.HandleFunc("/read/", read)
	http.ListenAndServe(":8081", nil)
}

func read(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprint(w, cookie)
}

func write(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:       "my-cookie",
		Value:      "test",
	})

	http.SetCookie(w, &http.Cookie{
		Name:       "my-cookie-1",
		Value:      "test",
	})
	fmt.Fprint(w, "Cookie has been set")
}
