package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", nil)
}

func read(w http.ResponseWriter, req *http.Request) {
	var count int
	c, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:  "my-cookie",
			Value: "1",
		})
	} else {
		count, _ = strconv.Atoi(c.Value)
		http.SetCookie(w, &http.Cookie{
			Name:  "my-cookie",
			Value: strconv.Itoa(count + 1),
		})
	}

	fmt.Fprint(w, "Visited times, "+strconv.Itoa(count))
}
