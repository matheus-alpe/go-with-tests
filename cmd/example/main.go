package main

import (
	"go-with-tests/internal/mocking"
	"os"
	"time"
)

func main() {
	sleeper := mocking.NewConfigurableSleeper(500*time.Millisecond, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
