package main

import "fmt"

/*
- Using COMPOSITE LITERAL:
	-create a SLICE of TYPE int
	- assign 10 VALUES to it
- Range over the slice and print the values out.
- Using format printing
	- print out the TYPE of the array
*/

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

}
