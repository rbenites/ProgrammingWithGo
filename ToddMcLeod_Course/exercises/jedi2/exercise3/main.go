package main

import "fmt"

/*
- create TYPED and UNTYPED constants
- Print the values of the constants
*/

const (
	a     = 42
	b int = 42
)

func main() {

	fmt.Print(a)
	fmt.Println(b)

}
