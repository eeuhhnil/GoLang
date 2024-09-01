//go:build !solution

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- time.Duration) {
	start := time.Now()

	response, err := http.Get(url)

	if err == nil {
		defer response.Body.Close()
	}
	res := time.Since(start)
	ch <- res
}

func main() {
	start := time.Now()

	ch := make(chan time.Duration)
	urls := os.Args[1:]
	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
	close(ch)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
