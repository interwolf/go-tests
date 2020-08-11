package racer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("normally returns the fast url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got, err := Racer(slowURL, fastURL)
		want := fastURL

		if err != nil {
			t.Fatalf("get an unexpected error: %v", err)
		}

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("returns an error if no response in 1 sec", func(t *testing.T) {
		slowServer := makeDelayedServer(3 * time.Second)
		fastServer := makeDelayedServer(2 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_, err := Racer(slowURL, fastURL)

		if err == nil {
			t.Error("expect an error but not")
		}

		fmt.Println(err)
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
