package racer

import (
	"fmt"
	"net/http"
	"time"
)

var oneSecTimeout = 1 * time.Second

// Racer returns the first responding url, timeout = 1 sec
func Racer(url1, url2 string) (winner string, err error) {
	return ConfigurableRacer(url1, url2, oneSecTimeout)
}

// ConfigurableRacer returns the first responding url
func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, err error) {

	select {
	case <-ping(url1):
		// fmt.Printf("result1: %v\n", result1)
		return url1, nil
	case <-ping(url2):
		// fmt.Printf("result2: %v\n", result2)
		return url2, nil
	case <-time.After(timeout):
		// fmt.Printf("result3: %v\n", result3)
		return "", fmt.Errorf("time out waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	// fmt.Printf("ch: %v\n", ch)

	go func() {
		http.Get(url)
		fmt.Printf("will close ch: %v\n", ch)
		close(ch)
	}()

	fmt.Printf("ch: %v\n", ch)
	return ch
}

// func Racer(url1, url2 string) (winner string) {

// 	duration1 := measureTime(url1)
// 	duration2 := measureTime(url2)

// 	fmt.Printf("duration1: %q, duration2: %q\n", duration1, duration2)

// 	if duration1 < duration2 {
// 		return url1
// 	}

// 	return url2
// }

func measureTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
