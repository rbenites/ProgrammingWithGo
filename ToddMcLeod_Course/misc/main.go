package main

import "fmt"

//we create VALUES of a certain type that are stored in VARIABLES
//those VARIABLES have IDENTIFIERS
var x int

type person struct {
	first string
	last  string
}

type foo int

var y foo

const bar int = 42

func main() {
	y = 42
	fmt.Printf("%T\n", int(y))
	fmt.Printf("%T\n", bar)
	fmt.Println(bar)

}
