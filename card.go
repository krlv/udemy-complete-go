package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Rank is a numeric representation of a card's rank
type Rank int

// Rank constants
const (
	_ Rank = iota
	_
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

// ToString returns string representation of a card's rank
func (r Rank) ToString() string {
	return [...]string{
		"", "", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King", "Ace",
	}[r]
}

// ToString returns string representation of a card's suite
func (s Suite) ToString() string {
	return [...]string{"Diamonds", "Hearts", "Clubs", "Spades"}[s]
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
