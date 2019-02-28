package main

import "fmt"

// ContactInfo struc
type ContactInfo struct {
	Email string
	Zip   int
}

// Person struct
type Person struct {
	FirstName string
	LastName  string
	Contacts  ContactInfo
}

// SetName for a person
func (p *Person) SetName(s string) {
	p.FirstName = s
}

func main() {
	thomas := Person{
		FirstName: "Thomas",
		LastName:  "Anderson",
		Contacts: ContactInfo{
			Email: "neo@matrix.net",
			Zip:   20001,
		},
	}

	// Thomas Anderson is Neo!
	thomas.SetName("Neo")

	fmt.Printf("%+v\n", thomas)
}
