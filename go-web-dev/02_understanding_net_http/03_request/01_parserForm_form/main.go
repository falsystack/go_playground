package main

import (
	"html/template"
	"log"
	"net/http"
)

/*
Form, PostFormを呼び出すためには先にParseFormを呼び出す必要がある。
*/

var tpl *template.Template

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("02_understanding_net_http/03_request/01_parserForm_form/index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
