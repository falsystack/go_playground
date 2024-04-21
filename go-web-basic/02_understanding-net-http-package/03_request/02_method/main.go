package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type cat int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (c cat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		Method:      req.Method,
		Submissions: req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	var c cat
	http.ListenAndServe(":8080", c)
}
