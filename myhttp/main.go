package myhttp

import (
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; encoding=utf8")
	io.WriteString(w, "<h1> Hello from index </h1>")
}

func Dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; encoding=utf8")
	io.WriteString(w, "<h1> Bark from a dog </h1>")
}

func Me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; encoding=utf8")
	io.WriteString(w, "<h1> Tarlan Huseynov </h1>")
}
