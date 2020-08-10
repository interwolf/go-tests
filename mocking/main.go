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

// ConfigurableSleeper allows specify duration and sleep func
type ConfigurableSleeper struct {
	duration  time.Duration
	sleepFunc func(time.Duration)
}

// Sleep of ConfigurableSleeper sleeps duratin seconds
func (c *ConfigurableSleeper) Sleep() {
	c.sleepFunc(c.duration)
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
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
