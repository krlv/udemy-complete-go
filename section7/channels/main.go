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

	for _, url := range urls {
		status := "not ok"
		if CheckURL(url) {
			status = "ok"
		}

		fmt.Printf("%s site is %s\n", url, status)
	}
}

// CheckURL returns true if URL is ok (1XX, 2XX, 3XX status code),
// otherwise returns false
func CheckURL(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}

	if resp.StatusCode > 399 {
		return false
	}

	return true
}
