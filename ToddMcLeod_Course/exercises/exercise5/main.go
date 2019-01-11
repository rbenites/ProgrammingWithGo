package main

import "fmt"

/*
BUILDING ON THE EX4 code
- at the package level scope, using the var keyword, create a VARIABLE with the IDENTIFIER "y"
	-the variable should be of the UNDERLYING type of x (int)
- in func main use CONVERSION to convert the TYPE of the VALUE stored in x to the UNDERLYING TYPE
	-Then use the short declaration to ASSIGN that value to y
- print out the value stored in y
-format print the type of y

*/

//new type called starwars that is of type int
type starWars int

//var x is of type starWars (which is of type int, so x is of type int as well)
var x starWars
var y int

func main() {

	//print out value
	fmt.Println(x)
	//print out type
	fmt.Printf("%T\n", x)
	//assign 42 to x
	x = 42
	//print the value of x
	fmt.Println(x)
	//assign the conversion of x from starWars to int type
	y = int(x)
	//print value sored in y
	fmt.Println(y)
	//print out type of y
	fmt.Printf("%T", y)

}
