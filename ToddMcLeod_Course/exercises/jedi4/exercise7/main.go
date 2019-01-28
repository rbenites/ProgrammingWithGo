package main

import "fmt"

/*
- Create a slice of a slice of string
	-ex: [][]string
	-Store the following
		-"James","Bond","Shaken, not stirred"
		-"Miss","Moneypenny","Helo, james"
- Range over the records, then range over the data in each record
*/

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Helo, james"}
	fmt.Println(mp)

	ss := [][]string{jb, mp}
	fmt.Println(ss)

	for i, s := range ss {
		fmt.Println("Record: ", i)
		for j, val := range s {
			fmt.Printf("\t index position: %v\t value: \t %v\n", j, val)
		}
	}
}
