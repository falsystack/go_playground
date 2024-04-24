package main

import (
	"go-web-tucker/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", myapp.NewHttpHandler())
}
