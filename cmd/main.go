package main

import (
	"go-with-tests/internal/mocking"
	"os"
)

func main() {
	mocking.Countdown(os.Stdout, &mocking.DefaultSleeper{})
}
