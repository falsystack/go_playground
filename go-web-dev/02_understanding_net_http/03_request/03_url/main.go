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
		Submissions url.Values
	}{
		req.Method,
		req.URL,
		req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("02_understanding_net_http/03_request/03_url/index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
