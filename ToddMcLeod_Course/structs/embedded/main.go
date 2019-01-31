package main

import "fmt"

type person struct {
	//field names. Must be unique
	//Identifier and type
	first string
	last  string
	age   int
}

type secretAgent struct {
	//anon type
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
}
