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

func NewDeck() Deck {
	var deck Deck

	ranks := [...]Rank{
		Two,
		Three,
		Four,
		Five,
		Six,
		Seven,
		Eight,
		Nine,
		Ten,
		Jack,
		Queen,
		King,
		Ace,
	}

	suites := []Suite{
		Diamonds,
		Hearts,
		Clubs,
		Spades,
	}

	for i, _ := range ranks {
		for j, _ := range suites {
			deck = append(deck, Card{ranks[i], suites[j]})
		}
	}

	return deck
}

func (d Deck) Print() {
	for _, card := range d {
		card.Print()
	}
}