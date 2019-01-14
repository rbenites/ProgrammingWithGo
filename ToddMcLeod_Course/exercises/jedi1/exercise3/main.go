package main

import "fmt"

/*
using the code for EX 2 in func main
- use fmt.Sprintf to print all of the VALUES to one single string.
- ASSIGN the returned value oy TYPE string using the shortcut declaration operator to a VARIABLE with the IDENTIFIER "s"
- print out the value stored by the variable "s"
*/

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// %v prints the value in the default format
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
