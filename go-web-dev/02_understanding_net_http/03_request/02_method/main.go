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
		Method     string
		Submission url.Values
	}{
		req.Method,
		req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("02_understanding_net_http/03_request/02_method/index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
