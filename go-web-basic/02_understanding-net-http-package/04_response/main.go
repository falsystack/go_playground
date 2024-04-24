package main

import (
	"fmt"
	"net/http"
)

type shower int

func (s shower) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("falsystack", "This is custom header.")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>This is Title</h1>")
}

func main() {
	http.ListenAndServe("localhost:8080", shower(0))
}
