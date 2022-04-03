package myhttp

import (
	"io"
	"net/http"
	"text/template"
)

var Tpl template.Template

func Index(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-type", "text/html; encoding=utf8")
	Tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
func Me(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-type", "text/html; encoding=utf8")
	r.FormValue("fname")
	Tpl.ExecuteTemplate(w, "me.gohtml", r.Form["fname"][0])
}

func Dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; encoding=utf8")
	io.WriteString(w, "<h1> Bark from a dog </h1>")
}

func init() {
	template.Must(Tpl.ParseGlob("myhttp/templates/*"))
}
