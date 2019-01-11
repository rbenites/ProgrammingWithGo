package main

import "fmt"

func main() {
	//n is the variable for what is to be printed.
	//err is the variable to hold the error response. if you dont want to print this you can replace with _
	n, err := fmt.Println("Hello, playground", 42, true)
	fmt.Println(n, err)

	/*statements:
	smallest element of a program that expresses some action to be carried out
	*/
	//x is declared and assigned a value
	x := 42
	fmt.Println(x)
	x = 100
	fmt.Println(x)
	//statement. Take up one line and made up of expressions.
	/*Expressions:
	combination of one or more explicit values, constans, variables, operators, and functions that the programming language interprets and computes to produce another value.
	*/
	y := 100 + 24
	fmt.Println(y)
}
