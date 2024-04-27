# Go web With Tucker

## testify
assertのみget
```shell
go get github.com/stretchr/testify/assert
```
結構見やすい！

### POST, GET以外のメソッドのテスト

`POST`, `GET` メソッド以外は基本提供されてない

```go
func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodDelete, ts.URL+"/users/1", nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}
```


## gorilla/mux
Package gorilla/mux is a powerful HTTP router and URL matcher for building Go web servers with 🦍
```shell
go get github.com/gorilla/mux
```

