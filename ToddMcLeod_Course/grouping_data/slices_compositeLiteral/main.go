package main

//a slice allows you to group together VALUES of the same type

import "fmt"

func main() {
	//coposite literal with values of the same type
	//x:=type{values}

	//slice of int assigned to x with values of 4,5,7,82,42
	x := []int{4, 5, 7, 82, 42}
	fmt.Println(x)
}
