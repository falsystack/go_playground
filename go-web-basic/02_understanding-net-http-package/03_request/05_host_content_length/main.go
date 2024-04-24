package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type shower int

func (s shower) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

func init() {
	tpl = template.Must(template.ParseFiles("02_understanding-net-http-package/03_request/05_host_content_length/index.gohtml"))
}

func main() {
	http.ListenAndServe(":8080", shower(0))
}
