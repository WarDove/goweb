package main

import (
	"github.com/WarDove/goweb/myhttp"
	. "net/http"
)

func main() {

	Handle("/", HandlerFunc(myhttp.Index))
	Handle("/me", HandlerFunc(myhttp.Me))
	Handle("/dog", HandlerFunc(myhttp.Dog))

	ListenAndServe(":8080", nil)

}
