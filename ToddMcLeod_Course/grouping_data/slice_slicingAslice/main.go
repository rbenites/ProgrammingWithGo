package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	//print whats in the 1 position
	fmt.Println(x[1])
	//print whats in x
	fmt.Println(x)
	//print from position 1 to the end
	fmt.Println(x[1:])
	//print from position 1 up to but not including position 3
	fmt.Println(x[1:3])

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
