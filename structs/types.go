package main

import "fmt"

// New custom type definition
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	// Go automatically dereferences p to access the firstName field of the person struct.
	p.firstName = newFirstName
}
