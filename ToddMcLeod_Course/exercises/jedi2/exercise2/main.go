package main

import "fmt"

/*
- Use the following operators to write expressions and assign their value to variables
-- g ==, h<=, i>=, j!=, k<, l>
*/

func main() {
	a := (42 == 42)
	fmt.Println(a)
	b := (1 <= 10)
	fmt.Println(b)
	c := (10 >= 1)
	fmt.Println(c)
	d := (100 != 1)
	fmt.Println(d)
	e := (1 < 10)
	fmt.Println(e)
	f := (10 > 1)
	fmt.Println(f)
}
