package app

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-web-tucker/todo/model"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestTodos(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(MakeHandler())
	defer ts.Close()

	resp, err := http.PostForm(ts.URL+"/todos", url.Values{
		"name": {"Test todo"},
	})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	var todo model.Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo")

	id1 := todo.ID

	resp, err = http.PostForm(ts.URL+"/todos", url.Values{
		"name": {"Test todo2"},
	})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo2")

	id2 := todo.ID

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	todos := []*model.Todo{}
	json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 2)

	for _, t := range todos {
		if t.ID == id1 {
			assert.Equal("Test todo", t.Name)
		} else if t.ID == id2 {
			assert.Equal("Test todo2", t.Name)
		} else {
			assert.Error(fmt.Errorf("testID should be id1 or id2"))
		}
	}

	resp, err = http.Get(ts.URL + "/complete-todo/" + strconv.Itoa(id1) + "?completed=true")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	todos = []*model.Todo{}
	json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 2)

	for _, t := range todos {
		if t.ID == id1 {
			assert.True(t.Completed)
		}
	}

	// GET, POSTしかサポートしないのでDELETEは直接リクエストを作る必要がある
	req, _ := http.NewRequest("DELETE", ts.URL+"/todos/"+strconv.Itoa(id1), nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	todos = []*model.Todo{}
	json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 1)
	assert.Equal(todos[0].ID, id2)

	// Refactoring 19:00
}
