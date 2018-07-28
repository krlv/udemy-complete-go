package main

import "testing"

func TestNewDeckFromFile(t *testing.T) {
	t.Skip()
}

func TestNewDeck(t *testing.T) {
	t.Skip()
}

func TestDeckShuffle(t *testing.T) {
	t.Skip()
}

func TestDeckDeal(t *testing.T) {
	t.Skip()
}

func TestDeckString(t *testing.T) {
	t.Skip()
}

func TestDeckSaveToFile(t *testing.T) {
	t.Skip()
}

func ExampleDeck_Print() {
	d := Deck{
		Card{Ace, Spades},
		Card{Queen, Hearts},
	}
	d.Print()
	// Output:
	// Ace of Spades
	// Queen of Hearts
}
