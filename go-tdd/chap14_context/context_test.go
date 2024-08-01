package chap14_context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	data := "hello, world"

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		svr.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Errorf("got %s, want %s", res.Body.String(), data)
		}

		//store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelld", func(t *testing.T) {
		store := &SpyStore{
			response: data,
			t:        t,
		}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel) // その関数が5ミリ秒で呼び出されるようにスケジュール
		req = req.WithContext(cancellingCtx)

		//res := httptest.NewRecorder()
		res := &SpyResponseWriter{}

		svr.ServeHTTP(res, req)

		if res.written {
			t.Errorf("a response should not have been written")
		}
	})

}
