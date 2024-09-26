package main

import "fmt"

func main() {
	integer_array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range integer_array {
		if num%2 == 0 {
			fmt.Printf("Num %d is even", num)
		} else {
			fmt.Printf("Num %d is odd", num)
		}
	}
}
