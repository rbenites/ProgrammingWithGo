package main

import "fmt"

/*
- create your own TYPE of init
- create a VARIABLE of the new TYPE with the IDENTIFIER "x" using VAR keyword
- in func main print out the value of x and then the type of x
- assign 42 to the VARIBALE x using = OPERATOR
- print out the value of x

*/

//new type called starwars that is of type int
type starWars int

//var x is of type starWars (which is of type int, so x is of type int as well)
var x starWars

func main() {

	//print out value
	fmt.Println(x)
	//print out type
	fmt.Printf("%T\n", x)
	//assign 42 to x
	x = 42
	//print the value of x
	fmt.Println(x)
}
