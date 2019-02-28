package main

import "strings"

// Hand is a slice of cards, similar to Deck
type Hand []Card

// String representation of a hand
func (h Hand) String() string {
	var strs []string

	for _, c := range h {
		strs = append(strs, c.String())
	}

	return strings.Join(strs, "\n")
}

// Print cards in a hand
func (h Hand) Print() {
	for _, card := range h {
		card.Print()
	}
}
