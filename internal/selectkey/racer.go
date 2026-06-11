package selectkey

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondsTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondsTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

type PingChan chan struct{}

func ping(url string) PingChan {
	ch := make(PingChan)
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
