package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"go-web-tucker/myupload"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUpload(t *testing.T) {
	assert := assert.New(t)
	path := ""
	file, _ := os.Open(path)
	defer file.Close()

	// byte 데이터를 생성
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	multi, err := w.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)
	io.Copy(multi, file)
	w.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-Type", w.FormDataContentType())

	myupload.UploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	// upload 한 파일과 오리지널 파일이 같은지 확인
	uploadFilePath := "uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath)
	assert.NoError(err)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(originData, uploadData)
}
