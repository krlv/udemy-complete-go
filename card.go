package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Rank is a numeric representation of a card's rank
type Rank int

// Rank constants
const (
	_ Rank = iota
	Joker
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

var ranks = []Rank{
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

// NewRank creates a card's rank from string representation
func NewRank(s string) Rank {
	switch strings.Title(s) {
	case "Two":
		return Two
	case "Three":
		return Three
	case "Four":
		return Four
	case "Five":
		return Five
	case "Six":
		return Six
	case "Seven":
		return Seven
	case "Eight":
		return Eight
	case "Nine":
		return Nine
	case "Ten":
		return Ten
	case "Jack":
		return Jack
	case "Queen":
		return Queen
	case "King":
		return King
	case "Ace":
		return Ace
	}

	// unhandled error
	return Ace
}

// ToString returns string representation of a card's rank
func (r Rank) ToString() string {
	return [...]string{
		"", "", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King", "Ace",
	}[r]
}

// Suite is a numeric representation of a card's suit
type Suite int

// Suite constants
const (
	Diamonds Suite = iota
	Hearts
	Clubs
	Spades
)

var suites = []Suite{
	Diamonds,
	Hearts,
	Clubs,
	Spades,
}

// ToString returns string representation of a card's suite
func (s Suite) ToString() string {
	return [...]string{"Diamonds", "Hearts", "Clubs", "Spades"}[s]
}

// SuiteFromString converts string representation to card's suite
func SuiteFromString(s string) Suite {
	smap := map[string]Suite{
		"Diamonds": Diamonds, "Hearts": Hearts,
		"Clubs": Clubs, "Spades": Spades,
	}

	return smap[s]
}

// Card type represents card's rank adn suite
type Card struct {
	Rank  Rank
	Suite Suite
}

// Print card's rank and suite
func (c Card) Print() {
	fmt.Println(c.Rank.ToString() + " of " + c.Suite.ToString())
}

// ToString returns string representation of a card
func (c Card) ToString() string {
	return c.Rank.ToString() + " " + c.Suite.ToString()
}

// NewCard generates new random card
func (d Deck) NewCard() Card {
	rand.Seed(time.Now().UnixNano())

	return Card{
		Rank:  ranks[rand.Intn(int(Ace))],
		Suite: suites[rand.Intn(int(Spades))],
	}
}
