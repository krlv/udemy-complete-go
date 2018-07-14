package main

import "testing"

func TestNewRank(t *testing.T) {
	var rank Rank

	rank = NewRank("two")
	if rank != Two {
		t.Errorf("Expected to create rank Two for value \"two\"")
	}

	rank = NewRank("three")
	if rank != Three {
		t.Errorf("Expected to create rank Three for value \"three\"")
	}

	rank = NewRank("four")
	if rank != Four {
		t.Errorf("Expected to create rank Four for value \"four\"")
	}

	rank = NewRank("Five")
	if rank != Five {
		t.Errorf("Expected to create rank Five for value \"Five\"")
	}

	rank = NewRank("Six")
	if rank != Six {
		t.Errorf("Expected to create rank Three for value \"Three\"")
	}

	rank = NewRank("Seven")
	if rank != Seven {
		t.Errorf("Expected to create rank Seven for value \"Seven\"")
	}

	rank = NewRank("Eight")
	if rank != Eight {
		t.Errorf("Expected to create rank Eight for value \"Eight\"")
	}

	rank = NewRank("Nine")
	if rank != Nine {
		t.Errorf("Expected to create rank Nine for value \"Nine\"")
	}

	rank = NewRank("Ten")
	if rank != Ten {
		t.Errorf("Expected to create rank Ten for value \"Ten\"")
	}

	rank = NewRank("Jack")
	if rank != Jack {
		t.Errorf("Expected to create rank Jack for value \"Jack\"")
	}

	rank = NewRank("Queen")
	if rank != Queen {
		t.Errorf("Expected to create rank Queen for value \"Queen\"")
	}

	rank = NewRank("King")
	if rank != King {
		t.Errorf("Expected to create rank King for value \"King\"")
	}

	rank = NewRank("Ace")
	if rank != Ace {
		t.Errorf("Expected to create rank Ace for value \"Ace\"")
	}

	rank = NewRank("none")
	if rank != Ace {
		t.Errorf("Expected to create rank Ace for unkown value")
	}
}
