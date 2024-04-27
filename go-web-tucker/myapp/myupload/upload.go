package myupload

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadsHandler(w http.ResponseWriter, req *http.Request) {
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
