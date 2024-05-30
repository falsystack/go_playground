package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := "/Users/yun/Downloads/test.csv"
	file, _ := os.Open(path)
	defer file.Close()

	buf := &bytes.Buffer{}
	// io.Writerは bufferを作ってあげれば良い
	w := multipart.NewWriter(buf)
	multi, err := w.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)

	io.Copy(multi, file)
	w.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	// w.FormDataContentType()はboundaryがついている
	req.Header.Set("Content-Type", w.FormDataContentType())

	uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	// file infoが返ってくる
	stat, err := os.Stat(uploadFilePath)
	log.Println(stat)
	assert.NoError(err)

	// uploadしたファイルとoriginファイルが一緒なのかをチェック
	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	uploadFile.Read(originData)

	assert.Equal(uploadData, originData)
}
