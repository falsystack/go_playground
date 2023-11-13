package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		id := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure: true,
			HttpOnly: true, // jsで接近不可
			Path:     "/",
		}
		http.SetCookie(w, c)
	}
	fmt.Println(c)
}
