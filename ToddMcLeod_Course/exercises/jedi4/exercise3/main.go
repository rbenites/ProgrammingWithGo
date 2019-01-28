package main

import "fmt"

/*
- Using the code from EX2
	- use SLICING to create the following new slices
	[42 43 44 45 46]
	[47 48 49 50 51]
	[44 45 46 47 48]
	[43 44 45 46 47]
	- Print it out
*/

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

}
