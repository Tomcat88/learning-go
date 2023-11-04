package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyResponseWriter struct {
	written bool
}

func (w *SpyResponseWriter) Header() http.Header {
	w.written = true
	return nil
}

func (w *SpyResponseWriter) Write([]byte) (int, error) {
	w.written = true
	return 0, errors.New("not implemented")
}

func (w *SpyResponseWriter) WriteHeader(statuscode int)  {
	w.written = true
}

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertStoreWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not cancelled")
	}
}

func (s *SpyStore) assertStoreWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was cancelled")
	}
}

func TestServer(t *testing.T) {
	t.Run("returns data from the store", func(t *testing.T) {
		data := "hello, world"
		store := SpyStore{response: data, t: t}
		srv := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}

		store.assertStoreWasNotCancelled()
	})
	t.Run("tells the store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := SpyStore{response: data, t: t}
		srv := Server(&store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		srv.ServeHTTP(response, request)

		if response.written {
			t.Error("Response should not be written")
		}
	})
}
