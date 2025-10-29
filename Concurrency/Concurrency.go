// This is a sample Go file for concurrency demonstration
// Web pinger
package main

import (
	"fmt"
	"net/http"
	"time"

	// Package de datos de resultado
	"GO_tutorial/Concurrency/Result"
)

func main() {
	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://stackoverflow.com",
		"https://google.com",
		"https://.com"}

	result := make(chan Result.Result)

	for _, url := range urls {
		// Send an HTTP GET request concurrently
		// send a get request
		// mesure how lo it takes (ms)
		// return throug a channel
		// Go routine
		go func(url string) {
			start := time.Now()

			_, err := http.Get(url)

			duration := time.Since(start)

			result <- Result.Result{
				URL:      url,
				Duration: duration,
				Error:    err,
			}
		}(url)
	}
	// Collect results
	for range urls {
		res := <-result
		if res.Error != nil {
			fmt.Printf(" %s -> error: %v\n", res.URL, res.Error)
		} else {
			fmt.Printf(" %s -> took: %v\n", res.URL, res.Duration)
		}
	}
	fmt.Println("Finish all pings")
}
