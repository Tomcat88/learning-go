package concurrency

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrTimeout = errors.New("timeout")

const defaultTimeout = 10 * time.Second

func WebsiteRacer(website1, website2 string) (string, error) {
	return ConfigurableRacer(website1, website2, defaultTimeout)
}

func ConfigurableRacer(website1, website2 string, timeout time.Duration) (string, error) {
	fmt.Println("here")
	select {
	case <-ping(website1):
		return website1, nil
	case <-ping(website2):
		return website2, nil
	case <-time.After(timeout):
		return "", ErrTimeout
	}
}

func measureURLLatency(url string) time.Duration {
	startA := time.Now()
	http.Get(url)
	return time.Since(startA)
}

func ping(url string) chan struct{} {
	r := make(chan struct{})
	go func() {
		http.Get(url)
		close(r)
	}()
	return r
}
