package main

import "fmt"

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range ints {
		s := "odd"

		if i%2 == 0 {
			s = "even"
		}

		fmt.Printf("%v is %s\n", i, s)
	}
}
