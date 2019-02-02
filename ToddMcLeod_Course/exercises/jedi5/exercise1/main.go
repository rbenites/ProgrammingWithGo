package main

import "fmt"

/*
- Create your own type "person" which will have an underlying type of "struct" so that it can store the following data:
	-first name
	-last name
	- favorite ice cream flavors
- Create two VALUES of TYPE person.
- Print out the values ranging over the elements in the slice which stores the fav flavors
*/

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "bond",
		favIceCream: []string{
			"Rum Raisin",
			"Cherry",
			"pistachio"},
	}
	p2 := person{
		first: "miss",
		last:  "moneypenny",
		favIceCream: []string{
			"chocolate", "double'o'fudge",
			"caramel"},
	}
	fmt.Println(p1.first, p1.last)
	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first, p2.last)
	for i, v := range p2.favIceCream {
		fmt.Println(i, v)
	}

}
