package main

import "net/http"

type home struct{}

func (h home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

type page struct {
	body string
}

func (p page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// echo back the page first URI
	w.Write([]byte(p.body))
}

// use map instead
type multiplexer map[string]http.Handler

func (m multiplexer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := m[r.RequestURI]; ok {
		handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

var mux = multiplexer{
	"/":            home{},
	"/references/": page{"references"},
	"/tutorials/":  page{"tutorials"},
}

func Index(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello i am index without subdomain"))

}

func main() {
	//http.HandleFunc("/", Index)
	//http.ListenAndServe(":8080", nil)                 // port 8080 is optional
	http.ListenAndServe("api.example.com:8080/", mux) // port 8080 is optional
}
