package myapp

import (
	"net/http/httptest"
	"testing"
)

func TestIndexPathHandler(t *testing.T) {
	httptest.NewRecorder()               // mock response
	httptest.NewRequest("GET", "/", nil) // mock request
}
