package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//new thing called deck that acts like a slice of string. This alloes us to tie any function to type deck
type deck []string

//wehn we call newDeck we create a new deck of type deck and return it by placing deck after the function name
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

//(d deck) = receiver.
//d= var d is a reference to the cards array. similar to this
//deck=every var of type deck can call this function on itself
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//deal function that takes in d deck as an argument and returns handSize as type int
//the size starts by taking taking from begining to imputed handsize and then from handsize to the end of the deck 
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//function with receiver that joins the deck into a string seperated by commas
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	//source equals a call to rand.NewSource that provides a new source seed generated from the time function and unix function
	source := rand.NewSource(time.Now().UnixNano())
	//r holds the New rand from source
	r := rand.New(source)

	for i := range d {
		//r of type rand has a function intn. Call this function and pass in an int and return a new int that s passed to newposition
		newPosition := r.Intn(len(d) - 1)
		//take what is at d-newposition and assign to d-i and whats at d-i and assign to d-newposition
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
