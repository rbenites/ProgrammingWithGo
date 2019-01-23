package main
/*
in general it is not advised to use arrays in GO. 
Instead use slices.
*/

import "fmt"

func main() {
	//an array of type int with 5 "spots"
	var x [5]int
	fmt.Println(x)
	//position 3 = 42
	x[3] = 42
	fmt.Println(x)

}
