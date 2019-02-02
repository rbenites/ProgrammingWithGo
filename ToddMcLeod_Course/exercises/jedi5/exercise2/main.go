package main

import "fmt"

/*
- Take the previous code
	- store the values of type person in a map with the key of last name
- Access each value in the map
- Print out the values, ranging over the slice
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

	//[key]value
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v.favIceCream {
			fmt.Println(i, v2)
		}
	}
}
