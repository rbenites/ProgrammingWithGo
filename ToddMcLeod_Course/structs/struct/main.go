package main

import "fmt"

/*
- A data structure
- Allows us to compse together values of different types
	- aggregaes data of diferent value types
*/

//composite or aggregate or complex data structure
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
	//prints the value of p1
	fmt.Println(p1)
	fmt.Println(p2)
//printe the palue of p1.first and p1.last
	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}
