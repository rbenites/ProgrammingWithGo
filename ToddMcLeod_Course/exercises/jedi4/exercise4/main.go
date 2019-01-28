package main

import "fmt"

/*
- start with the following slice:
	- x:=[]int{42,43,44,45,46,47,48,49,50,51}
- append to the the slice the following value:
	- 52
- in ONE STATEMENT append to the slice the following values:
	- 53
	- 54
	- 55
- print out the slice
- append to the slice the following slice
	- x:=[]int{56,57,58,59,60}
print out the slice
*/

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}
