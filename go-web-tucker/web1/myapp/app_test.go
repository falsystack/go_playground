package myapp

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"

	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	indexHandler(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World, ", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)
	name := "test"

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name="+name, nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal(fmt.Sprintf("Hello %s, %s", name, name), string(data))
}

func TestHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("GET", "/foo", nil)
	res := httptest.NewRecorder()

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestHandler_WithoJson(t *testing.T) {
	assert := assert.New(t)

	// buffer -> io.Reader 로 변함
	r := strings.NewReader(`{"first_name":"test", "last_name":"key", "email":"test@test.com"}`)

	req := httptest.NewRequest("POST", "/foo", r)
	res := httptest.NewRecorder()

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	user := &User{}
	assert.Equal(http.StatusCreated, res.Code)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("key", user.LastName)
	assert.Equal("test", user.FirstName)
	assert.Equal("test@test.com", user.Email)
}
