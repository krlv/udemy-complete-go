package main

// Deck is a slice of cards
type Deck []Card

// Hand is also a slice of cards
type Hand []Card

// Print cards in a hand
func (h Hand) Print() {
	for _, card := range h {
		card.Print()
	}
}

// NewDeck generates new deck of cards
func NewDeck() Deck {
	var deck Deck

	for i := range ranks {
		for j := range suites {
			deck = append(deck, Card{ranks[i], suites[j]})
		}
	}

	return deck
}

// Print cards in the deck
func (d Deck) Print() {
	for _, card := range d {
		card.Print()
	}
}
