package main

import "fmt"

/*
- assign an int to a variable
- print the int in decimal, binary, and hex
- shift the bits of the int over 1 position to the left and assin that to a variable
- print the int in decimal, binary, and hex
*/

func main() {
	a := 2
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)

}
