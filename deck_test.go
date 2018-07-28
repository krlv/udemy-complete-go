package main

import "testing"

func TestNewDeckFromFile(t *testing.T) {
	t.Skip()
}

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck of 52 cards, got %d cards", len(d))
	}

	if d[0] != (Card{Two, Diamonds}) {
		t.Errorf("Expected first card to be Two of Diamonds, got %d", d[0])
	}

	if d[51] != (Card{Ace, Spades}) {
		t.Errorf("Expected first card to be Ace of Spades, got %d", d[0])
	}
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
