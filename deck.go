package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Deck is a slice of cards
type Deck []Card

// NewDeckFromFile restore a deck from a file
func NewDeckFromFile(filename string) (Deck, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cards []Card
	for _, v := range strings.Split(string(bytes), "\n") {
		card := strings.Split(v, " ")

		cards = append(cards, Card{
			NewRank(card[0]),
			NewSuite(card[1]),
		})
	}

	return Deck(cards), nil
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

// Shuffle a deck
func (d Deck) Shuffle() {
	src := rand.NewSource(time.Now().UTC().UnixNano())
	rnd := rand.New(src)

	rnd.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
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
