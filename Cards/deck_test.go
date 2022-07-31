package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// act
	result := newDeck()

	// assert
	if len(result) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(result))
	}
	if result[0] != "Ace of Clubs" {
		t.Errorf("Expected Ace of Clubs, but got %s", result[0])
	}
	if result[len(result)-1] != "King of Spades" {
		t.Errorf("Expected King of Spades, but got %s", result[len(result)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// arrange
	filename := "_decktesting"
	os.Remove(filename)
	d := newDeck()

	// act
	d.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)

	// assert
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(loadedDeck))
	}

	// cleanup
	os.Remove(filename)
}
