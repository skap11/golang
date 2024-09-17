package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colors)
	printMap(colors)

	// Ways of initializing an empty map
	// var asciiNumber map[string]string
	asciiNumber := make(map[string]int)
	asciiNumber["A"] = 97
	asciiNumber["B"] = 98

	// Ignores if the key is not there in the map
	delete(asciiNumber, "B")

	fmt.Println(asciiNumber)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex for " + color + " is " + hex)
	}
}
