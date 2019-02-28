package main

import "fmt"

// Bot interface
type Bot interface {
	GetGreeting(string) string
}

// EnglishBot structure
type EnglishBot struct {
	Greeting string
}

// GetGreeting for name s in English
func (eb EnglishBot) GetGreeting(s string) string {
	return fmt.Sprintf("🇬🇧 "+eb.Greeting, s)
}

// SpanishBot structure
type SpanishBot struct {
	Greeting string
}

// GetGreeting for name s in Spanish
func (sb SpanishBot) GetGreeting(s string) string {
	return fmt.Sprintf("🇪🇸 "+sb.Greeting, s)
}

// PrintGreeting using language bot
func PrintGreeting(b Bot, s string) {
	fmt.Println(b.GetGreeting(s))
}

func main() {
	eb := EnglishBot{"Hello, %s!"}
	sb := SpanishBot{"Hola, %s!"}

	PrintGreeting(eb, "Neo")
	PrintGreeting(sb, "Neo")
}
