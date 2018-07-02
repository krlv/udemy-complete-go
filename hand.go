package main

import "strings"

// Hand is a slice of cards, similar to Deck
type Hand []Card

// Print cards in a hand
func (h Hand) Print() {
	for _, card := range h {
		card.Print()
	}
}

// ToString returns string representation of a hand
func (h Hand) ToString() string {
	var strs []string

	for _, c := range h {
		strs = append(strs, c.ToString())
	}

	return strings.Join(strs, "\n")
}
