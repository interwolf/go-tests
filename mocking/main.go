package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Countdown prints 3\n 2\n 1\n go!\n
func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(writer, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(writer, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
