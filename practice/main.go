package main

import "fmt"

func main() {
	var s string

	ints := make([]int, 11)
	for i := range ints {
		ints[i] = i

		s = "odd"
		if i%2 == 0 {
			s = "even"
		}

		fmt.Printf("%v is %s\n", i, s)
	}
}
