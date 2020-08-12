package contextv1

import (
	"context"
	"fmt"
	"net/http"
)

// Store is an interface for fetching a string
type Store interface {
	Fetch(ctx context.Context) (string, error)
	// Cancel()
}

// ServerHanderFunc returns a HandlerFunc writing a string fetching from Store
func ServerHanderFunc(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprint(w, data)
	}
}
