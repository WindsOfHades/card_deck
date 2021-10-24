package main

func main() {
    // cards := newDeck()
    // cards.print()
    // hand, remainingDeck := deal(cards, 5)
    // remainingDeck.print()
    // hand.saveToFile("hand.txt")
    /////////////////////////////////////////
    // readDeck := newDeckFromFile("hand.txt")
    // readDeck.print()
    // readDeck.shuffle()
    // readDeck.print()

    cards := newDeckFromFile("struct-string.txt")
    cards.print()
}
