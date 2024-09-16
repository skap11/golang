package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got: %v", len(d))
	}
}

func TestNewDeck_FirstCard(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got: %v", d[0])
	}
}

func TestNewDeck_LastCard(t *testing.T) {
	d := newDeck()
	lastElement := len(d) - 1

	if d[lastElement] != "Four of Clubs" {
		t.Errorf("Expected first card to be 'Four of Clubs', but got: %v", d[lastElement])
	}
}

func TestReadDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := readDeckFromFile("_decktesting")

	if len(loadedDeck) != 11 {
		t.Errorf("Expected deck length of 16, but got: %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
