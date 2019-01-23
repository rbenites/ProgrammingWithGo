package main

import "fmt"

/*
- create a for loop using the following syntax:
	for {}
- have it print out the years you have been alive
*/

func main() {
	i := 1980
	for {
		if i > 2019 {
			break
		}
		fmt.Println(i)
		i++
	}
}
