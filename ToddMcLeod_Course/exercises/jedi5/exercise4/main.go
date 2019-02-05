package main

import "fmt"

/*
- Create and use an anon struct
*/

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"water",
		},
	}
	fmt.Println(s)

	for k, v := range s.friends {
		fmt.Println(k, v)

	}

	for i,val:=range s.favDrinks{
		fmt.Println(i, val)
	}
}
