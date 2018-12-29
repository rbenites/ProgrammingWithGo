package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//function(shorthand receiverType) functionName() returnType

func (eb englishBot) getGreeting() string {
	//VERY custom logic
	return "Hi there"
}

func (sb spanishBot) getGreeting() string {
	//VERY custom logic
	return "Hola"
}
