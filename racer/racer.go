package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	resp, err := http.Get(a)
	if err == nil {
		resp.Body.Close()
	}
	aDuration := time.Since(startA)

	startB := time.Now()
	resp, err = http.Get(b)
	if err == nil {
		resp.Body.Close()
	}
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}
	return b
}
