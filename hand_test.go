package main

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
