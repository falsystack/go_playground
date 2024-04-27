package main

import (
	"go-web-tucker/webs/webapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", webapp.NewHandler())

}
