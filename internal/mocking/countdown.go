package mocking

import (
	"fmt"
	"io"
	"iter"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func NewConfigurableSleeper(
	duration time.Duration,
	sleep func(time.Duration),
) *ConfigurableSleeper {
	return &ConfigurableSleeper{
		duration: duration,
		sleep:    sleep,
	}
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := range countDownFrom(3) {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, finalWord)
}
