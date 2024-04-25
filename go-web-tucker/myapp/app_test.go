package myapp

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexPathHandler(t *testing.T) {
	// Go言語はテストもSimpleだな！
	assert := assert.New(t)

	resp := httptest.NewRecorder()                 // mock response
	req := httptest.NewRequest("GET", "/bar", nil) // mock request

	mux := NewHttpHandler()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code)
	data, _ := io.ReadAll(resp.Body)
	// 検証Fail時に見えるメッセージがが見やすい！！いいね！
	assert.Equal("Hello World", string(data))
}

func TestWithParam(t *testing.T) {
	name := "Naver"

	assert := assert.New(t)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", fmt.Sprintf("/bar?name=%s", name), nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code)
	data, _ := io.ReadAll(resp.Body)
	assert.Equal(fmt.Sprintf("Hello %s", name), string(data))
}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusBadRequest, resp.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", strings.NewReader(`{"first_name":"kakao","last_name":"daum","email":"kakao@daum.net"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusCreated, resp.Code)

	u := new(User)
	err := json.NewDecoder(resp.Body).Decode(u)
	assert.Nil(err)
	assert.Equal("kakao", u.FirstName)
	assert.Equal("daum", u.LastName)
	assert.Equal("kakao@daum.net", u.Email)
}
