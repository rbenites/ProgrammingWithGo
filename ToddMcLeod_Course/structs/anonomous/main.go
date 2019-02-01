package main

import "fmt"

//named struct
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "james",
		last:  "bond",
		age:   32,
	}
	fmt.Println(p1)
	//anonomous struct
	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "Miss",
		last:  "Moneypenny",
		age:   32,
	}
	fmt.Println(p2)
}
