package main

import "testing"

func TestPersonSetName(t *testing.T) {
	var p Person

	(&p).SetName("Sam")

	if p.FirstName != "Sam" {
		t.Errorf("Expected name \"Sam\", got \"%s\"", p.FirstName)
	}
}

func Example_main() {
	main()
	// Output:
	// {FirstName:Neo LastName:Anderson Contacts:{Email:neo@matrix.net Zip:20001}}
}
