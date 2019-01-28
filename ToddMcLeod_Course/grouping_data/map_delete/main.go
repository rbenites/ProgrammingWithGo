package main

import "fmt"

func main() {
	//map[type]type{ key:value,}
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	//delete an entry
	delete(m, "James")
	fmt.Println(m)
}
