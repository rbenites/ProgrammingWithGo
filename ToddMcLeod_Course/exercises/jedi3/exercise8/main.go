package main

import "fmt"

/*- create a program that uses a switch statement with no switch expression specified
 */

func main() {
	switch {
	case false:
		fmt.Println("Should not print")
	case true:
		fmt.Println("Should print")

	}

}
