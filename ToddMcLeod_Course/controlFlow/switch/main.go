package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Should not print")
	case (2 == 4):
		fmt.Println("Should not print")
	case (3 == 3):
		fmt.Println("This Prints")
		//this allows the code to keep going to the next check and print
		//DONT USE FALLTHROUGH!!!!!
		fallthrough
	case (4 == 4):
		fmt.Println("Should Print")
		fallthrough
	case (7 == 9):
		fmt.Println("Not true1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
		//this prints if nothing is true
	default:
		fmt.Println("this is default")
	}

}
