package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("print 3 to 1 then Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &CountdownSleeperSpy{}

		Countdown(buffer, spy)
		fmt.Printf("CallStrings: %v\n", spy.CallStrings)

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("speel-write interleaved", func(t *testing.T) {
		spy := &CountdownSleeperSpy{}
		Countdown(spy, spy)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(spy.CallStrings, want) {
			t.Errorf("got: %v, want: %v", spy.CallStrings, want)
		}

	})

}

func TestConfigurableSleeper(t *testing.T) {
	toSleep := 5 * time.Second
	spy := &SleepTimeSpy{}
	sleeper := ConfigurableSleeper{toSleep, spy.Sleep}
	sleeper.Sleep()

	if spy.slept != toSleep {
		t.Errorf("expect slept: %v, but got: %v", toSleep, spy.slept)
	}

}

const write = "write"
const sleep = "sleep"

// CountdownSleeperSpy mocks Countdown
type CountdownSleeperSpy struct {
	CallStrings []string
}

//Sleep spies by appending "sleep"
func (s *CountdownSleeperSpy) Sleep() {
	s.CallStrings = append(s.CallStrings, sleep)
}

func (s *CountdownSleeperSpy) Write(p []byte) (n int, err error) {
	s.CallStrings = append(s.CallStrings, write)
	return
}

type SleepTimeSpy struct {
	slept time.Duration
}

func (s *SleepTimeSpy) Sleep(duration time.Duration) {
	s.slept = duration
}

// // SpySleeper records how many calls to its Sleep()
// type SpySleeper struct {
// 	NumCalls int
// }

// // Sleep of SpySleeper
// func (s *SpySleeper) Sleep() {
// 	s.NumCalls++
// }
