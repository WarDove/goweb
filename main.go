package main

import (
	"github.com/WarDove/goweb/myhttp"
	"net/http"
)

func main() {

	http.HandleFunc("/", myhttp.Index)
	http.HandleFunc("/me", myhttp.Me)
	http.HandleFunc("/dog", myhttp.Dog)

	http.ListenAndServe(":8080", nil)

}
