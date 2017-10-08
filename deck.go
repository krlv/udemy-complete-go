package main

import (
	"fmt"
	"strings"
	"time"
	"math/rand"
)

// Type Hand is a slice of cards
type Hand []Card

// Type Deck is also a slice of cards
type Deck []Card

// Type Card ...
type Card struct {
	Rank string
	Suit string
}

func (c Card) Print () {
	fmt.Println(strings.Title(c.Rank) + " of " + strings.Title(c.Suit))
}

func (h Hand) Print () {
	for _, card := range h {
		card.Print()
	}
}

func (d Deck) NewCard () Card {
	ranks := []string {
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"jack",
		"queen",
		"king",
		"ace",
	}

	suits := []string {
		"diamonds",
		"hearts",
		"clubs",
		"spades",
	}

	rand.Seed(time.Now().UnixNano())

	return Card{
		Rank: ranks[rand.Intn(len(ranks))],
		Suit: suits[rand.Intn(len(suits))],
	}
}