package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type card struct {
    rank string
    suite string
}
type deck []card

func newDeck() deck {
    cards := deck{}

    suites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
    values := []string{"Ace", "King", "Queen", "Jack", "10", "9", "8", "7", "6", "5", "4", "3", "2"}

    for _, suite:= range suites {
        for _, value:= range values {
            cards = append(cards, card {
                suite: suite,
                rank: value,
            })
        }
    }

    return cards
}

func (c card) toString() string {
    return c.rank + "-" + c.suite
}

func (d deck) toString() string{
    var temp []string
    for _, value:= range d {
        temp = append(temp, value.toString())
    }
    return strings.Join(temp, ",")
}

func (d deck) print() {
    for i, card:= range d {
        fmt.Println(i+1, card)
        fmt.Println("-----")
    }
}

func deal(d deck, handsize int) (deck, deck) {
    return d[:handsize], d[handsize:]
}

func (d deck) shuffle() {
    for i := range d {
        rand.Seed(time.Now().UnixNano())
        randomNumber := rand.Intn(len(d))
        fmt.Println(randomNumber)
        d[i], d[randomNumber] = d[randomNumber], d[i]
    }
}

func (d deck) saveToFile(filename string ) error{
    return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}


func newDeckFromFile(filename string) deck {
    d := deck{}
    byteSlice, err := ioutil.ReadFile(filename)
    if (err != nil) {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }
    slices := strings.Split(string(byteSlice), ",")
    for _, value := range slices {
        cardString := strings.Split(value, "-")
        d = append(d, card {
            rank: cardString[0],
            suite: cardString[1],
        })
    }
    return d
}
