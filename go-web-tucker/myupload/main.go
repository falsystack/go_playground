package myupload

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("myupload/public")))
	http.HandleFunc("/uploads", UploadsHandler)
	http.ListenAndServe(":8080", nil)
}

// 20:03까지함
