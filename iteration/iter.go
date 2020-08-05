package iteration

// const count = 5

// Repeat repeates c five times
func Repeat(c string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += c
	}
	return repeated
}
