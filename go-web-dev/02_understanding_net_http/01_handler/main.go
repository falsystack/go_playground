package main

import (
	"fmt"
	"net/http"
)

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("何でも書ける")
}

func main() {
}
