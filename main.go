package main

func main() {
	// full deck
	cards := NewDeck()
	cards.Print()

	// hand of 5 cards
	hand := cards.Deal(5)
	hand.Print()

	// deck without 5 cards
	cards.Print()

	// save deck to a file
	cards.SaveToFile("test.txt")

	// restore deck from a file
	cards, _ = NewDeckFromFile("test.txt")
	cards.Print()
}
