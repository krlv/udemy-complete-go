package main

// Type Hand is a slice of cards
type Hand []Card

// Type Deck is also a slice of cards
type Deck []Card

func (h Hand) Print() {
	for _, card := range h {
		card.Print()
	}
}
