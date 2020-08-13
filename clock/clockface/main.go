package main

import (
	"os"
	"time"

	"github.com/interwolf/go-tests/clockface"
)

func main() {
	// fmt.Printf("Hello, Clock!\n")
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
