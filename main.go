package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Seed the randomizer
	rand.Seed(time.Now().UTC().UnixNano())

	// full deck
	cards := NewDeck()
	cards.Shuffle()
	cards.Print()

	// hand of 5 cards
	hand := cards.Deal(5)
	hand.Print()

	// deck without 5 cards
	cards.Print()

	// save deck to a file
	cards.SaveToFile("test.txt")

	// restore deck from a file
	cards, err := NewDeckFromFile("test.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	cards.Print()
}
