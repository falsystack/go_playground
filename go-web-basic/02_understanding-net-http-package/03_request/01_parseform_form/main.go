package main

import (
	"html/template"
	"log"
	"net/http"
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
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)

}

func main() {
	var c cat
	http.ListenAndServe(":8080", c)
}
