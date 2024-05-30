package myapp

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	// mock server
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))

}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	// mock server
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Contains(string(data), "Get UserInfo by")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	// mock server
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Contains(string(data), "User ID:89")
}
