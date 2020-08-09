package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet prints Hello to name
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// GreetHandler writes Hello world in response
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "John")
	http.ListenAndServe(":5000", http.HandlerFunc(GreetHandler))
}
