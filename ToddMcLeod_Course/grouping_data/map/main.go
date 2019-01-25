package main

import "fmt"

/*
map is a key value pair

*/

func main() {
	//map[type]type{ key:value,}
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barney"])
	//check to see if the value is in the map
	v, ok := m["Barney"]
	fmt.Println(v)
	fmt.Println(ok)
	//the below will not print since barney would be not ok
	if v, ok := m["Barney"]; ok {
		fmt.Println("This is the IF Print", v)
	}
	//This will print since James in the map
	if v, ok := m["James"]; ok {
		fmt.Println("This is the IF Print", v)
	}
}