package main

import "fmt"

func main() {
	colors := map[string]int{
		"black": 0x000000,
		"red":   0xff0000,
		"green": 0x00ff00,
		"blue":  0x0000ff,
	}

	colors["white"] = 0xffffff
	delete(colors, "black")

	for c, h := range colors {
		fmt.Printf("HEX representation for %s is %x\n", c, h)
	}
}
