package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Creates and returns a list of playing cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Loops through the deck
// and prints the value
// of every single card
// We use a reciever
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deals a fixed ammount of cards
// and returns the hand and the
// remaining deck of cards
// as two separate types
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Converts a deck into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Saves a deck to a file on the hard drive
func (d deck) saveToFile(filname string) error {
	return ioutil.WriteFile(filname, []byte(d.toString()), 0666)
}

// Reading a new deck from a file
func newDeckFromFile(filename string) deck {
	biteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - Log the error and return a call to newDeck
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(biteSlice), ",")
	// type conversion
	return deck(s)
}

// Shuffles the deck of cards
func (d deck) shuffle() {
	// Making a logic that generates
	// a new pseudo-random number
	// each time by generating a Rand type
	// and feeding the time as a seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
