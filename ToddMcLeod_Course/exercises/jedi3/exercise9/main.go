package main

import "fmt"

/*
-Create a program that uses a switch statement with the switch expression specified
-expression = Type string with the IDENTIFIER "favSport"
*/

func main() {
	favSport := "Football"
	switch favSport {
	case "skiing":
		fmt.Println("Go to the Mountains")
	case "surfing":
		fmt.Println("Go to the beach")
	case "Baseball":
		fmt.Println("Go to the field and play BALL")
	case "Football":
		fmt.Println("FOOOOTBAAAALLLLLL!!!!")
	default:
		fmt.Println("Can you really call that a sport?")
	}
}
