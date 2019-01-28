package main

import "fmt"

/*
- Using the following slice:
	- x:=[]int{42,43,44,45,46,47,48,49,50,51}
	use APPEND and SLICING to get the following values:
		-[42, 43, 44, 45, 46, 48, 49, 50]
	- print it out

*/

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:5], x[6:9]...)
	fmt.Println(x)
}
