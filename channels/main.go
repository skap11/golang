package main

import "fmt"

func main() {
	// websites := []string{"https://google.com", "https://facebook.com"}

	// wait for child routines to finish
	// channel - Meant for sharing data(typed) b/w go routines
	// c := make(chan string)

	// for _, website := range websites {
	// 	// checkLink(website)
	// 	go checkLink(website, c)
	// }

	c := make(chan string)
	for i := 0; i < 4; i++ {
		go printString("Hello there!", c)
	}

	for s := range c {
		fmt.Println(s)
	}

	// for {
	// 	fmt.Println(<-c)
	// }
}

func printString(s string, c chan string) {
	fmt.Println(s)
	c <- "Done printing."
}

// func checkLink(link string, c chan string) {
// 	_, error := http.Get(link)
// 	if error != nil {
// 		fmt.Println(link, "might be down!!")
// 		return
// 	}

// 	fmt.Println(link, "is up!!!")
// }
