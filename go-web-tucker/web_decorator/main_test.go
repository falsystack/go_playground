package main

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"go-web-tucker/web_decorator/myapp"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexPage(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(myapp.NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Equal("Hello World!", string(data))
}

func TestDecoHandler(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(myapp.NewHandler())
	defer ts.Close()

	// 基本loggerのアウトプットはstd ioになっているから(terminalに出力する)
	//　binary bufferに向きを変えてテストで確認できるようにする
	buf := bytes.Buffer{}
	log.SetOutput(&buf)

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	r := bufio.NewReader(&buf)
	line, _, err := r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Started")

	line, _, err = r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Completed")
}
