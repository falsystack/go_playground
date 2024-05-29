package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("web1/web4-1/public")))
	http.HandleFunc("/uploads", uploadsHandler)

	http.ListenAndServe(":3000", nil)
}

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile, header, err := r.FormFile("upload_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	dirname := "uploads"
	os.MkdirAll(dirname, os.ModePerm) // 0777 rwx
	filePath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filePath) // 빈 파일 생성
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	io.Copy(file, uploadFile) // form file -> 빈 file에 복사
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Uploaded successfully")
}
