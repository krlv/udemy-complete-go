package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
	}

	ch := make(chan func() (string, bool))

	go func() {
		for {
			for _, url := range urls {
				go CheckURL(url, ch)
			}

			time.Sleep(5 * time.Second)
		}
	}()

	for f := range ch {
		url, st := f()

		status := "not ok"
		if st {
			status = "ok"
		}

		fmt.Printf("%s site is %s\n", url, status)
	}
}

// CheckURL sends true to the channel if URL is ok (1XX, 2XX, 3XX status code),
// otherwise sends false
func CheckURL(url string, ch chan func() (string, bool)) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- getURLStatusClosure(url, false)
		return
	}

	if resp.StatusCode > 399 {
		ch <- getURLStatusClosure(url, false)
		return
	}

	ch <- getURLStatusClosure(url, true)
	return
}

func getURLStatusClosure(url string, status bool) func() (string, bool) {
	return func() (string, bool) {
		return url, status
	}
}
