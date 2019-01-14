package main

import "fmt"

//int = whole number
//float = real numbers (has deimals)

var x int
var y float64

func main() {

	x = 42
	y = 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

}
