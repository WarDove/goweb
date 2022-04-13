package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var visitcount int

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	w.Header()
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "visitcount: " + strconv.FormatInt(int64(visitcount), 10),
		Path:  "/"})

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
	visitcount++
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
