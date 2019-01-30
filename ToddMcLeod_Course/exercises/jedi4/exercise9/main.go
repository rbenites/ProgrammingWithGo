package main

import "fmt"

/*
- using the code from ex8
	-add a record to the map
- print the updated map out using range
*/

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

	m["fleming_ian"] = []string{"steaks", "cigards", "espionage"}

	//for k (key), v (value)= range of m (the map of key value pair)
	for k, v := range m {
		//print the key
		fmt.Println(k)
		//for each i (index), v2 is the range of v (the slice of string)
		for i, v2 := range v {
			i++
			//print a tab then the index then the value2
			fmt.Println("\t", i, v2)
		}
	}
}
