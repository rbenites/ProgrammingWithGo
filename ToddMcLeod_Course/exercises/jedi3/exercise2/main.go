package main

import "fmt"

/*
- Print every rune code point of the uppercase alphabet three times
EX:
1
	U+0041 'A'
	U+0041 'A'
	U+0041 'A'
*/

func main() {
	for o := 65; o <= 90; o++ {
		fmt.Println(o)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\n", o)
		}
	}
}
