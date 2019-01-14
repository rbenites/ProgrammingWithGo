package main

import "fmt"

func main() {
	//for statement with for clause (nested for loop)
	//for intit; condition; post {}
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}

	//for statement with single clause
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")

}
