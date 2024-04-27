package main

import (
	"go-web-tucker/web_decorator/myapp"
	"net/http"
)

func main() {
	mux := myapp.NewHandler()

	http.ListenAndServe(":8080", mux)
}
