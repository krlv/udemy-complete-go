package main

import "testing"

func TestHandString(t *testing.T) {
	h := Hand{
		Card{Ace, Spades},
		Card{Queen, Hearts},
	}
	s := `Ace Spades
Queen Hearts`

	if h.String() != s {
		t.Errorf("Expected \"%s\", got \"%s\"", s, h)
	}
}

func ExampleHand_Print() {
	h := Hand{
		Card{Ace, Spades},
		Card{Queen, Hearts},
	}
	h.Print()
	// Output:
	// Ace of Spades
	// Queen of Hearts
}
