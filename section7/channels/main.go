package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
	}

	ch := make(chan bool)

	for _, url := range urls {
		go CheckURL(url, ch)
	}

	for _, url := range urls {
		status := "not ok"

		if <-ch {
			status = "ok"
		}

		fmt.Printf("%s site is %s\n", url, status)
	}
}

// CheckURL sends true to the channel if URL is ok (1XX, 2XX, 3XX status code),
// otherwise sends false
func CheckURL(url string, ch chan bool) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- false
		return
	}

	if resp.StatusCode > 399 {
		ch <- false
		return
	}

	ch <- true
	return
}
