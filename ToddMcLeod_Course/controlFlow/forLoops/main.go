package main

import "fmt"

func main() {
	//for intit; condition; post {}
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	//for statement with single clause
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")

}
