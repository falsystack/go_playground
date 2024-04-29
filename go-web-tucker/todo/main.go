package main

import (
	"github.com/urfave/negroni"
	"go-web-tucker/todo/app"
	"log"
	"net/http"
)

func main() {

	m := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", n)
	if err != nil {
		panic(err)
	}
}
