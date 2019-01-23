package main

import "fmt"

/*
- create a program that cuses "else if" and "else"
*/

func main() {
	x := "James Bernd"
	if x == "James Bond" {
		println("Bond", x)
	} else if x != "James Bond" {
		fmt.Println("You're not 007!")
	}

}
