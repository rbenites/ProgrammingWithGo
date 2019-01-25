package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	//append returns a another slice
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	//append x with all the values from y
	x = append(x, y...)
	fmt.Println(x)

	//x=append and take all the values from begining up to 2nd position and append from the 4th position to the end
	//7 and 8 are removed
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}