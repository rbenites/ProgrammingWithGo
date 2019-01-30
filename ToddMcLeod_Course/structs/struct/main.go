package main

import "fmt"

/*
- A data structure
- Allows us to compse together values of different types
	- aggregaes data of diferent value types
*/
type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
