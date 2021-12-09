package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type bulgarianBot struct{}

func main() {
	eb := englishBot{}
	bb := bulgarianBot{}

	printGreeting(eb)
	printGreeting(bb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// when we don`t plan on using the value of the reciever
// "eb" we can omit it and the complier won`t throw an error`
func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	// non reusable!
	return "Hi There!"
}

func (bulgarianBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	// non reusable!
	return "Zdravei!"
}
