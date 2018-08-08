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

func main() {
	thomas := Person{
		FirstName: "Thomas",
		LastName:  "Anderson",
		Contacts: ContactInfo{
			Email: "neo@matrix.net",
			Zip:   20001,
		},
	}
	fmt.Printf("%+v\n", thomas)
}
