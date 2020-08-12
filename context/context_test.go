package contextv1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	// cancelled bool
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	dataChannel := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				fmt.Println("Got cancelled!")
				return
			default:
				time.Sleep(1 * time.Millisecond)
				result += string(c)
			}
		}
		dataChannel <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-dataChannel:
		return res, nil
	}

	// return s.response, nil
}

type SypResponseWriter struct {
	written bool
}

func (s *SypResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SypResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SypResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestHandler(t *testing.T) {
	data := "Hello, Yiming!"
	// cancelled := false

	t.Run("Cancel work if told so", func(t *testing.T) {
		store := &SpyStore{data}
		handler := ServerHanderFunc(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellableCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellableCtx)
		response := &SypResponseWriter{}

		handler.ServeHTTP(response, request)

		if response.written {
			t.Errorf("Cancel in time and thus should not write response")
		}
	})

	t.Run("Fetch data if not cancelled", func(t *testing.T) {
		store := &SpyStore{data}
		handler := ServerHanderFunc(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got: %s, want: %s", response.Body.String(), data)
		}

		// if store.cancelled {
		// 	t.Errorf("store should not be told cancelled but is!")
		// }
	})

}
