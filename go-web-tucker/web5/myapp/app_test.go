package myapp

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
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

func TestCreateUser(t *testing.T) {
	// given
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	body := `{"first_name": "ky", "last_name":"Yun", "email":"test@test.com"}`

	// when
	resp, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(body))

	// then
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := &User{}
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.Equal("2", user.ID)

	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(user.ID))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}
