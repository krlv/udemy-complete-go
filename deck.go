package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Type Rank is a numeric representation of a card's rank
type Rank int

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

func (r Rank) String() string {
	return [...]string{
		"", "", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King", "Ace",
	}[r]
}

// Type Suite is a numeric representation of a card's suit
type Suite int

const (
	Diamonds Suite = iota
	Hearts
	Clubs
	Spades
)

func (s Suite) String() string {
	return [...]string{"Diamonds", "Hearts", "Clubs", "Spades"}[s]
}

// Type Hand is a slice of cards
type Hand []Card

// Type Deck is also a slice of cards
type Deck []Card

// Type Card ...
type Card struct {
	Rank  Rank
	Suite Suite
}

func (c Card) Print() {
	fmt.Println(c.Rank.String() + " of " + c.Suite.String())
}

func (h Hand) Print() {
	for _, card := range h {
		card.Print()
	}
}

func (d Deck) NewCard() Card {
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

	rand.Seed(time.Now().UnixNano())

	return Card{
		Rank:  ranks[rand.Intn(int(Ace))],
		Suite: suites[rand.Intn(int(Spades))],
	}
}
