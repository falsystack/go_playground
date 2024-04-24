package main

import (
	"io"
	"net/http"
)

type cruiser int

func (c cruiser) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Cruiser ServeHTTP!")
}

type land int

func (l land) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Land ServeHTTP!")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/land", land(0))
	mux.Handle("/cruiser", cruiser(0))

	http.ListenAndServe(":8080", mux)
}
