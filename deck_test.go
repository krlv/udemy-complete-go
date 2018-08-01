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
		t.Errorf("Expected the first card to be Two of Diamonds, got %d", d[0])
	}

	if d[51] != (Card{Ace, Spades}) {
		t.Errorf("Expected the last card to be Ace of Spades, got %d", d[0])
	}
}

func TestDeckShuffle(t *testing.T) {
	d1 := NewDeck()
	d2 := NewDeck()
	d2.Shuffle()

	if d1[0] == d2[0] {
		t.Errorf("Expected the first card to be different for shuffled deck")
	}

	if d1[51] == d2[51] {
		t.Errorf("Expected the last card to be different for shuffled deck")
	}
}

func TestDeckDeal(t *testing.T) {
	d := NewDeck()
	h := d.Deal(5)

	if len(d) != 47 {
		t.Errorf("Expected deck of 47 cards, got %d cards", len(d))
	}

	if len(h) != 5 {
		t.Errorf("Expected hand of 5 cards, got %d cards", len(h))
	}

	for _, dc := range d {
		for _, hc := range h {
			if hc == dc {
				t.Errorf("Expected deck to not contains cards from hand")
			}
		}
	}
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
