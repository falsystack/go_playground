package webapp

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

	// mock get request
	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// assert body data
	data, _ := io.ReadAll(resp.Body)
	assert.Equal("Hello World!", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	// mock server
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// mock get request
	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// assert body data
	data, _ := io.ReadAll(resp.Body)
	assert.Contains(string(data), "Get UserInfo by /users")
}

func TestGetUsers(t *testing.T) {
	assert := assert.New(t)

	// mock server
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// mock get request
	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// assert body data
	data, _ := io.ReadAll(resp.Body)
	assert.Contains(string(data), "User Id:89")

}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"kunwoong", "last_name":"yun", "email":"test@test.com"}`),
	)
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	u := new(User)

	err = json.NewDecoder(resp.Body).Decode(u)
	assert.NoError(err)
	assert.NotEqual(0, u.ID)

	id := u.ID
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	u2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(u2)
	assert.NoError(err)

	assert.Equal(u.ID, u2.ID)
	assert.Equal(u.FirstName, u2.FirstName)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodDelete, ts.URL+"/users/1", nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}
