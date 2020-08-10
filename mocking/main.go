package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper is for mocking system sleep
type Sleeper interface {
	Sleep()
}

// SpySleeper records how many calls to its Sleep()
type SpySleeper struct {
	NumCalls int
}

// Sleep of SpySleeper
func (s *SpySleeper) Sleep() {
	s.NumCalls++
}

// DefaultSleeper is the real sleeper
type DefaultSleeper struct{}

// Sleep of DefaultSleeper call time.Sleep
func (d *DefaultSleeper) Sleep() {
	time.Sleep(2 * time.Second)
}

// Countdown prints 3\n 2\n 1\n go!\n
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}

	sleeper.Sleep()
	fmt.Fprint(writer, "Go!")
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
