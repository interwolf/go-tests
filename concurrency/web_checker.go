package concurrency

// WebChecker checks a url
type WebChecker func(string) bool

type resultType struct {
	string
	bool
}

// CheckWebsites checks a slice of urls
func CheckWebsites(wc WebChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan resultType)

	for _, url := range urls {
		go func(u string) {
			// fmt.Printf("url: %q\n", u)
			resultChannel <- resultType{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	// time.Sleep(1 * time.Second)

	return results
}
