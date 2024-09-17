package main

import "fmt"

func main() {
	// Rely on the order of fields in the definition
	alex := person{"Alex", "Anderson", contactInfo{"alex.anderson@gmail.com", 123412}}
	palex := person{
		firstName: "Palex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex.anderson@gmail.com",
			zipCode: 123412,
		},
	}
	fmt.Println(alex, palex)

	var jalex person

	jalex.firstName = "Jalex"
	jalex.lastName = "Anderson"

	// Print struct fields along with value
	jalex.print()
	jalexPointer := &jalex
	jalexPointer.updateName("Alex")
	jalex.print()
}
