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

func TestDeckSaveToFile(t *testing.T) {
	t.Skip()
}

func TestDeckString(t *testing.T) {
	d := Deck{
		Card{Ace, Spades},
		Card{Queen, Hearts},
	}
	s := `Ace Spades
Queen Hearts`

	if d.String() != s {
		t.Errorf("Expected \"%s\", got \"%s\"", s, d)
	}
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
