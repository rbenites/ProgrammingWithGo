package main

/*
- Create a map with a key of TYPE string and a value of Type []string
	- this will be a person's "last_first" name
- this will store their favorite things
- store three records in the map
- print out all of the values along with their index position in the slice
*/
import (
	"fmt"
)

func main() {

	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	//print what is stored in m
	//fmt.Println(m)

	/*
		for key, v := range m {
			fmt.Println(key)
			fmt.Println(v)
		}
	*/

	//for k (key), v (value)= range of m (the map of key value pair)
	for k, v := range m {
		//print the key
		fmt.Println(k)
		//for each i (index), v2 is the range of v
		for i, v2 := range v {
			i++
			//print a tab then the index then the value2
			fmt.Println("\t", i, v2)
		}
	}
}
