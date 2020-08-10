package concurrency

import (
	"reflect"
	"testing"
)

func mockWebChecker(url string) bool {
	if url == "http://abcdzym.com/" || url == "http://wwwwzym.net/" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://nicexlab.com/",
		"http://abcdzym.com/",
		// "http://wwwwzym.net/",
		"https://www.google.com/",
	}

	want := map[string]bool{
		"https://www.google.com/": true,
		"http://nicexlab.com/":    true,
		"http://abcdzym.com/":     false,
		// "http://wwwwzym.net/": false,
	}

	got := CheckWebsites(mockWebChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}
