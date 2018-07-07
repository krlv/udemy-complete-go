package main

import (
	"io/ioutil"
	"strings"
)

// Deck is a slice of cards
type Deck []Card

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

// Deal a new hand with n cards from the deck
func (d *Deck) Deal(n int) Hand {
	var hand []Card
	deck := *d
	hand, *d = []Card(deck[:n]), []Card(deck[n:])
	return Hand(hand)
}

// Print cards in the deck
func (d Deck) Print() {
	for _, card := range d {
		card.Print()
	}
}

// ToString returns string representation of a deck
func (d Deck) ToString() string {
	var strs []string

	for _, c := range d {
		strs = append(strs, c.ToString())
	}

	return strings.Join(strs, "\n")
}

// SaveToFile saves a deck to a file
func (d Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0644)
}
