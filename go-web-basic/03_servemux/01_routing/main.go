package main

import (
	"io"
	"net/http"
)

type golang int

func (g golang) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/awesome":
		io.WriteString(w, "Golang is awesome")
	case "/good":
		io.WriteString(w, "Golang is good")
	}
}

func main() {
	http.ListenAndServe(":8080", golang(1))
}
