package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("upload/public")))
	http.HandleFunc("/uploads", uploadsHandler)
	http.ListenAndServe(":8080", nil)
}

func uploadsHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	uploadFile, h, err := req.FormFile("upload_file")
	defer uploadFile.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	dirname := "./uploads"
	os.MkdirAll(dirname, os.ModePerm)
	filepath := fmt.Sprintf("%s/%s", dirname, h.Filename)
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Uploaded successfully"+filepath)
}

// 20:03까지함
