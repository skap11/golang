package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{"https://google.com", "https://facebook.com"}

	for _, website := range websites {
		// checkLink(website)
		go checkLink(website)
	}

	// wait for child routines to finish
}

func checkLink(link string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, "might be down!!")
		return
	}

	fmt.Println(link, "is up!!!")
}
