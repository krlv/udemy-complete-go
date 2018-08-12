package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "https://httpbin.org/user-agent"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}
