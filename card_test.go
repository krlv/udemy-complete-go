package main

import "testing"

func TestNewRank(t *testing.T) {
	if NewRank("two") != Two {
		t.Errorf("Expected to create rank Two for value \"two\"")
	}

	if NewRank("three") != Three {
		t.Errorf("Expected to create rank Three for value \"three\"")
	}

	if NewRank("four") != Four {
		t.Errorf("Expected to create rank Four for value \"four\"")
	}

	if NewRank("Five") != Five {
		t.Errorf("Expected to create rank Five for value \"Five\"")
	}

	if NewRank("Six") != Six {
		t.Errorf("Expected to create rank Three for value \"Three\"")
	}

	if NewRank("Seven") != Seven {
		t.Errorf("Expected to create rank Seven for value \"Seven\"")
	}

	if NewRank("Eight") != Eight {
		t.Errorf("Expected to create rank Eight for value \"Eight\"")
	}

	if NewRank("Nine") != Nine {
		t.Errorf("Expected to create rank Nine for value \"Nine\"")
	}

	if NewRank("Ten") != Ten {
		t.Errorf("Expected to create rank Ten for value \"Ten\"")
	}

	if NewRank("Jack") != Jack {
		t.Errorf("Expected to create rank Jack for value \"Jack\"")
	}

	if NewRank("Queen") != Queen {
		t.Errorf("Expected to create rank Queen for value \"Queen\"")
	}

	if NewRank("King") != King {
		t.Errorf("Expected to create rank King for value \"King\"")
	}

	if NewRank("Ace") != Ace {
		t.Errorf("Expected to create rank Ace for value \"Ace\"")
	}

	if NewRank("none") != Ace {
		t.Errorf("Expected to create rank Ace for unkown value")
	}
}

func TestRankString(t *testing.T) {
	if Two.String() != "Two" {
		t.Errorf("Expected to return string \"Two\" for rank Two")
	}

	if Three.String() != "Three" {
		t.Errorf("Expected to return string \"Three\" for rank Three")
	}

	if Four.String() != "Four" {
		t.Errorf("Expected to return string \"Four\" for rank Four")
	}

	if Five.String() != "Five" {
		t.Errorf("Expected to return string \"Five\" for rank Five")
	}

	if Six.String() != "Six" {
		t.Errorf("Expected to return string \"Six\" for rank Six")
	}

	if Seven.String() != "Seven" {
		t.Errorf("Expected to return string \"Seven\" for rank Seven")
	}

	if Eight.String() != "Eight" {
		t.Errorf("Expected to return string \"Eight\" for rank Eight")
	}

	if Nine.String() != "Nine" {
		t.Errorf("Expected to return string \"Nine\" for rank Nine")
	}

	if Ten.String() != "Ten" {
		t.Errorf("Expected to return string \"Ten\" for rank Ten")
	}

	if Jack.String() != "Jack" {
		t.Errorf("Expected to return string \"Jack\" for rank Jack")
	}

	if Queen.String() != "Queen" {
		t.Errorf("Expected to return string \"Queen\" for rank Queen")
	}

	if King.String() != "King" {
		t.Errorf("Expected to return string \"King\" for rank King")
	}

	if Ace.String() != "Ace" {
		t.Errorf("Expected to return string \"Ace\" for rank Ace")
	}

	rank := Rank(100).String()
	if rank != "Not Defined" {
		t.Errorf("Expected to return string \"Not Defined\" for invalid rank")
	}
}

func TestNewSuite(t *testing.T) {
	if NewSuite("diamonds") != Diamonds {
		t.Errorf("Expected to create suite Diamonds for value \"diamonds\"")
	}

	if NewSuite("Hearts") != Hearts {
		t.Errorf("Expected to create suite Hearts for value \"Hearts\"")
	}

	if NewSuite("clubs") != Clubs {
		t.Errorf("Expected to create suite Clubs for value \"clubs\"")
	}

	if NewSuite("Spades") != Spades {
		t.Errorf("Expected to create suite Spades for value \"Spades\"")
	}

	if NewSuite("none") != Spades {
		t.Errorf("Expected to create suite Spades for unkown value")
	}
}

func TestSuiteString(t *testing.T) {
	if Diamonds.String() != "Diamonds" {
		t.Errorf("Expected to return string \"Diamonds\" for suite Diamonds")
	}

	if Hearts.String() != "Hearts" {
		t.Errorf("Expected to return string \"Hearts\" for suite Hearts")
	}

	if Clubs.String() != "Clubs" {
		t.Errorf("Expected to return string \"Clubs\" for suite Clubs")
	}

	if Spades.String() != "Spades" {
		t.Errorf("Expected to return string \"Spades\" for suite Spades")
	}

	suite := Suite(5).String()
	if suite != "Not Defined" {
		t.Errorf("Expected to return string \"Not Defined\" for invalid suite")
	}
}

func TestCardString(t *testing.T) {
	c := Card{Ace, Spades}
	if c.String() != "Ace Spades" {
		t.Errorf("Expected to return string \"Ace Spades\", got %s", c.String())
	}

	c = Card{Two, Diamonds}
	if c.String() != "Two Diamonds" {
		t.Errorf("Expected to return string \"Two Diamonds\", got %s", c.String())
	}
}

func ExampleCard_Print() {
	c := Card{Ace, Spades}
	c.Print()

	c = Card{Two, Diamonds}
	c.Print()
	// Output:
	// Ace of Spades
	// Two of Diamonds
}
