package main

// decSize := 20
// var decSize int

// func main() {
// 	// declares and assigns a value to a variable called card
// 	// var card string = "Ace of Spades"
// 	decSize = 20
// 	// card := "Ace of Spades" // New variable assignment
// 	fmt.Println(decSize)
// }

func main() {
	// cards := newDeck()
	// cards.saveToFile("mycards")
	cards := readDeckFromFile("my")
	cards.print()

}

// type color string

// func (c color) describe(description string) string {
// 	return string(c) + " " + description
// }
