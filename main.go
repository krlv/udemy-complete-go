package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func newCard() string {
	cards := []string {
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
	card := cards[rand.Intn(len(cards))];
	suit := suits[rand.Intn(len(suits))];

	return strings.Title(card) + " of " + suit;
}

func main() {
	cards := []string{newCard(), newCard()}

	for _, card := range cards {
		fmt.Println(card)
	}
}