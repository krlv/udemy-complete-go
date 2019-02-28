package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	n := os.Args[1]
	if n == "" {
		fmt.Fprint(os.Stderr, "Filename required!\n")
		os.Exit(1)
	}

	f, err := os.Open(n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
