package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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
		req.Method,
		req.URL,
		req.Form,
		req.Header,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("02_understanding_net_http/03_request/04_header/index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
