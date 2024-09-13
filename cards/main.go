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
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// printState()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
