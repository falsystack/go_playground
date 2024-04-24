package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type Reimen int

func (r Reimen) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions map[string][]string
		Header      http.Header
	}{
		Method:      req.Method,
		URL:         req.URL,
		Submissions: req.Form,
		Header:      req.Header,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("02_understanding-net-http-package/03_request/04_header/index.gohtml"))
}

func main() {
	var r Reimen
	http.ListenAndServe(":8080", r)
}
