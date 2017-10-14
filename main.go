package main

import (
	"fmt"
)

func main() {
	var deck Deck

	cards := Hand{
		deck.NewCard(),
		deck.NewCard(),
		deck.NewCard(),
	}
	cards.Print()

	// compare with built-in receiver's Print
	for _, card := range cards {
		fmt.Println(card)
	}
}
