package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type Ramen int

var tpl *template.Template

func (r Ramen) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
	}{
		Method:      req.Method,
		URL:         req.URL,
		Submissions: req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("02_understanding-net-http-package/03_request/03_url/index.gohtml"))
}

func main() {
	var r Ramen
	http.ListenAndServe(":8080", r)
}
