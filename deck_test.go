package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
    d := newDeck()
    if len(d) != 52 {
        t.Errorf("deck with wrong length")
    }
    if d[0].rank != "Ace" || d[0].suite != "Spades" {
        t.Errorf("Expected {Ace, Spades} - actual %s", d[0])
    }
    if d[len(d)-1].rank != "2" || d[len(d)-1].suite != "Clubs" {
        t.Errorf("Expected {2, Clubs} - actual %s", d[len(d)-1])
    }
}

func TestSaveToFile(t *testing.T) {
    cleanUpFile()

    newDeck := deck{ card {rank: "Ace", suite: "Spades"}}
    newDeck.saveToFile("_deckTesting")

    loadedDeck := newDeckFromFile("_deckTesting")
    if loadedDeck[0].rank != "Ace" || loadedDeck[0].suite != "Spades" {
        t.Errorf("Expected to save and load a deck with card: Ace-Spades got: %v", loadedDeck[0])
    }

    cleanUpFile()
}

func cleanUpFile() {
    os.Remove("_deckTesting")
}