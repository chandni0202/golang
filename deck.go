package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new deck of typeslice of string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"spades", "diamonds", "hearts", "clubs"}
	cardValuew := []string{"ace", "two", "three", "four"}

	for _, suit := range cardSuits {
		for _, values := range cardValuew {
			cards = append(cards, values+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		nePos := r.Intn(len(d) - 1)

		d[i], d[nePos] = d[nePos], d[i]
	}
}
