package main

import (
	"net/http"
	"html/template"
)

type HtmlSimple struct {
	Title 				string
}


func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./view/main.tmpl")
	t.Execute(w, nil)
}