package main

import "testing"

func TestEnglishBotGetGreeting(t *testing.T) {
	eb := EnglishBot{"Test %s!"}
	gr := eb.GetGreeting("test")

	xp := "🇬🇧 Test test!"
	if gr != xp {
		t.Errorf("Expected to have %s, got %s", xp, gr)

	}
}

func TestSpanishBotGetGreeting(t *testing.T) {
	sp := SpanishBot{"Test %s!"}
	gr := sp.GetGreeting("test")

	xp := "🇪🇸 Test test!"
	if gr != xp {
		t.Errorf("Expected to have %s, got %s", xp, gr)

	}
}

func Example_main() {
	main()
	// Output:
	// 🇬🇧 Hello, Neo!
	// 🇪🇸 Hola, Neo!
}
